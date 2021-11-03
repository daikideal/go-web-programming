package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// リクエストを解析してからFormフィールドにアクセスする必要がある
	r.ParseForm()
	fmt.Fprintln(w, r.Form) // => URLの値を取り扱う
	fmt.Fprintln(w, r.PostForm) // => フォームのkey:valueのみを取り扱う, application/x-www-form-urlencodedしかサポートしていない
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

// Goサーバを起動し、client.htmlをブラウザ上でローカルに開き、送信ボタンをクリックする
