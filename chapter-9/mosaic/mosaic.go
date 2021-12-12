package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"sync"
)

var TILESDB map[string][3]float64

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

// 1. 画像の平均色を求める
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}

	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// 2. 指定された幅に画像をリサイズする
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// 3. タイル画像データベースをメモリ内に作成
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannnot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

// 4. 最も値が近い画像を見つけ出す
// func nearest(target [3]float64, db *map[string][3]float64) string {
// 	var filename string
// 	smallest := 1000000.0
// 	for k, v := range *db {
// 		dist := distance(target, v)
// 		if dist < smallest {
// 			filename, smallest = k, dist
// 		}
// 	}
// 	delete(*db, filename)
// 	return filename
// }
func (db *DB) nearest(target [3]float64) string {
	var filename string
	// 1. ロックするとフラグmutexが設定される
	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename)
	// 2. アンロックするとフラグmutexが解除される
	db.mutex.Unlock()
	return filename
}

// 1. mapではなく構造体DBの参照を渡す
func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	// 2. チャネルを作成
	c := make(chan image.Image)
	sp := image.Point{0, 0}
	// 3. 無名のゴルーチンを作成
	go func() {
		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}
				// 4. DBのメソッドnearestを呼び出して最も値の近いタイルを得る
				nearest := db.nearest(color)
				file, err := os.Open(nearest)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error:", err)
					}
				} else {
					fmt.Println("error:", nearest)
				}
				file.Close()
			}
		}
		c <- newimage.SubImage(newimage.Rect)
	}()
	return c
}

// 画像をつなぎ合わせる
// `cut`から返されたチャネルを使用する
func combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	// 1. チャネルを作成
	c := make(chan string)

	// 2. 無名のゴルーチンを作成(ロジックのメイン部分)
	go func() {
		// 3. 元画像の全ての区画が最終的な画像にコピーされるまで待つ
		var wg sync.WaitGroup
		img := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			// 4. 元画像の区画がコピーされるたびにカウンタを減算する
			wg.Done()
		}
		// 5. WaitGroupのカウンタを4に設定
		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool
		// 6. ループして、終了を待ち続ける
		for {
			// 7. 第一のチャネルを選択
			select {
			case s1, ok1 = <-c1:
				go copy(img, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(img, s2.Bounds(), s2, image.Point{r.Max.X / 2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(img, s3.Bounds(), s3, image.Point{r.Min.X, r.Max.Y / 2})
			case s4, ok4 = <-c4:
				go copy(img, s4.Bounds(), s4, image.Point{r.Max.X / 2, r.Max.Y / 2})
			}
			// 8. 全てのチャネルが閉じられていればループを抜ける
			if ok1 && ok2 && ok3 && ok4 {
				break
			}
		}
		// 9. 全区画のコピーが終わるまで待つ
		wg.Wait()

		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}

// 5. 2点間のユークリッド距離を計算
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

// 6. 2条の計算
func sq(n float64) float64 {
	return n * n
}

// 7. モザイク写真を生成するたびにタイル画像データベースを複製
func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
