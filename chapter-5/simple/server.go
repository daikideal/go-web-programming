package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// テンプレートを解析する
	t, _ := template.ParseFiles("tmpl.html")
	// テンプレートにデータを当てはめる
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
