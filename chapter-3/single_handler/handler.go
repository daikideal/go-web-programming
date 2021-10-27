package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// 単一のハンドラ(通常の場合にはマルチプレクサを使用する)
// 全てのリクエストに対し"Hello World!"を返す
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
