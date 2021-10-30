package main

import (
	"fmt"
	"net/http"
)

// 既存のインタフェースが存在する、あるいはハンドラとしても使える型がほしい場合は、
// そのインタフェースにメソッド`ServeHTTP`を追加すればURLに割り当てられるハンドラを得られる。
// => Webアプリケーションのモジュール性を高められる

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	// ハンドラを直接定義
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// ハンドラを割り当て
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
