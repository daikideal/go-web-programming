package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 使用するテンプレートを全て解析する
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello, World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
