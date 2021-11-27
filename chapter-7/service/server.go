package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// 1. リクエストを正しい関数に振り分けるためのハンドラ関数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 2. 投稿の取り出し
//
// 検証:
// 	curl -i -X GET http://127.0.0.1:8080/post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 1. データベースから構造体Postにデータ取得
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// 2. 構造体PostをJSON文字列に組み換え
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// 3. JSONをResponseWriterに書き出し
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// 3. 投稿の作成
//
// 検証:
// 	curl -i -X POST -H "Content-Type: application/json" -d '{"content":"My first post", "author":"Sau Sheong"}' http://127.0.0.1:8080/post/
// 	psql -U gwp -d gwp -c 'select * from posts;'
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// 1. バイト列を作成
	body := make([]byte, len)
	// 2. バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var post Post
	// 3. バイト列を構造体Postに組み換え
	json.Unmarshal(body, &post)
	// 4. データベースのレコードを作成
	err = post.create()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

// 4. 投稿の更新
//
// 検証:
// 	curl -i -X PUT -H "Content-Type: application/json" -d '{"content":"Updated post", "author":"Sau Sheong"}' http://127.0.0.1:8080/post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 1. データベースから構造体Postにデータ取得
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	// 2. リクエスト本体からJSONデータを読み出し
	r.Body.Read(body)
	// 3. JSONデータを構造体Postに組み換え
	json.Unmarshal(body, &post)
	// 4. データベースを更新
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// 5. 投稿の削除
//
// 検証:
// 	curl -i -X DELETE http://127.0.0.1:8080/post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 1. データベースから構造体Postにデータを取得
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// 2. データベースから投稿データを削除
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
