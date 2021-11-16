package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// データの保存
func store(data interface{}, filename string) {
	// ゼロサイズのバッファにメモリを割り当てる(https://golang.org/doc/effective_go#allocation_new)
	// 実質的にはバイトデータの可変バッファであり、Read, Writeを持っている
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// データの読み込み
func load(data interface{}, filename string) {
	// ファイルから生データを読み出す
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// 生データからバッファを生成する
	// 実質的には生データをメソッドRead, Writeに渡している
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "post1")

	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
