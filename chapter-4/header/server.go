package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	// 全てのヘッダを取得
	h := r.Header
	// 特定のヘッダを取得
	// h := r.Header["Accept-Encoding"] // => 文字列のマップ
	// h := r.Header.Get("Accept-Encoding") // => カンマ区切りのリスト

	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
