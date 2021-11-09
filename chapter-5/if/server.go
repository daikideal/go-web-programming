package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().Unix())

	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
