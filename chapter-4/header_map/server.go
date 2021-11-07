package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`

	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Println(w, "そのようなサービスはありません。他を当たってください。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	// WriteHeaderは呼び出した直後にヘッダが変更されるため、
	// 先にLocationヘッダを追加しておく必要がある
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	server.ListenAndServe()
}

// curl -i 127.0.0.1:8080/redirect
// または、ブラウザで表示するとリダイレクトされることが確認できる
