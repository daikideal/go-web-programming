package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// MultipartForm.Fileを使用する場合
// func process(w http.ResponseWriter, r *http.Request) {
// 	r.ParseMultipartForm(1024) // まずparseする必要がある
// 	fileHeader := r.MultipartForm.File["uploaded"][0]
// 	file, err := fileHeader.Open()
// 	if err == nil {
// 		data, err := ioutil.ReadAll(file)
// 		if err == nil {
// 			fmt.Fprintln(w, string(data))
// 		}
// 	}
// }

// FormFileを使用する場合
func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
