package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	// bodyのメモリ上の最大容量は10MB
	r.ParseMultipartForm(10485760)
	// アップロードされたファイルとタイルサイズを取得
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	// アップロードされたターゲット画像をデコード
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	// タイル画像データベースを複製
	db := cloneTilesDB()

	// 1. 独立して処理できるように画像を分割して処理を分散(ファン・アウト)
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	// 1. 独立して処理できるように画像を分割して処理を分散(ファン・アウト)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	// 1. 独立して処理できるように画像を分割して処理を分散(ファン・アウト)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	// 1. 独立して処理できるように画像を分割して処理を分散(ファン・アウト)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)

	// 2. 画像をつなげて処理結果をまとめる(ファン・イン)
	c := combine(bounds, c1, c2, c3, c4)

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   <-c,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)

	// // 新規モザイク写真画像を作成
	// newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))

	// // 各タイルの基準ピクセルを設定
	// sp := image.Point{0, 0}
	// // ターゲット画像(の各区画)を巡回
	// for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
	// 	for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
	// 		r, g, b, _ := original.At(x, y).RGBA()
	// 		color := [3]float64{float64(r), float64(g), float64(b)}

	// 		nearest := nearest(color, &db)
	// 		file, err := os.Open(nearest)

	// 		if err == nil {
	// 			img, _, err := image.Decode(file)
	// 			if err == nil {
	// 				t := resize(img, tileSize)
	// 				tile := t.SubImage(t.Bounds())
	// 				tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)

	// 				// 得られたタイルを作成しておいたモザイク写真内に描画
	// 				draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
	// 			} else {
	// 				fmt.Println("error:", err, nearest)
	// 			}
	// 		} else {
	// 			fmt.Println("error:", err, nearest)
	// 		}
	// 		file.Close()
	// 	}
	// }

	// // 6. JPEGにエンコードし、base64文字列に変換してブラウザに送信
	// buf1 := new(bytes.Buffer)
	// jpeg.Encode(buf1, original, nil)

	// originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	// buf2 := new(bytes.Buffer)
	// jpeg.Encode(buf2, newimage, nil)

	// mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	// t1 := time.Now()
	// images := map[string]string{
	// 	"original": originalStr,
	// 	"mosaic":   mosaic,
	// 	"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	// }
	// t, _ := template.ParseFiles("results.html")
	// t.Execute(w, images)
}
