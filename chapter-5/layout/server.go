package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// func process(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("layout.html")
// 	// 第2引数で実行するテンプレートの名前を明示的に指定する
// 	t.ExecuteTemplate(w, "layout", "")
// }

// 同じ名前のテンプレートを切り替えて利用するハンドラ
// func process(w http.ResponseWriter, r *http.Request) {
// 	rand.Seed(time.Now().Unix())
// 	var t *template.Template

// 	if rand.Intn(10) > 5 {
// 		t, _ = template.ParseFiles("layout.html", "red_hello.html")
// 	} else {
// 		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
// 	}
// 	t.ExecuteTemplate(w, "layout", "")
// }

// ブロックアクションを使用
func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template

	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}

// curl -i 127.0.0.1:8080/process
