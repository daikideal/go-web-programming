package main

import (
	"fmt"
	"golang.org/x/net/http2" // go1.6以降では不要
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})

	// crypto/tls/generate_cert.goで生成した証明書と鍵を使用
	server.ListenAndServeTLS("cert.pem", "key.pem")
}

// サーバがHTTP/2で動作しているかどうかはcURLで確認する
// curl -I --http2 --insecure https://localhost:8080/
