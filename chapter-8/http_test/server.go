package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Db      *sql.DB // Postの全てのメソッドは構造体のフィールドの1つ`sql.DB`を使うようになる
	Id      int     `json:"id"`
	Content string  `json:"content"`
	Author  string  `json:"author"`
}

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

func main() {
	var err error
	db, err := sql.Open(
		"postgres",
		"user=gwp dbname=gwp password=password sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: ":8080",
	}
	// 1. 構造体Postを渡してhandleRequestを登録(間接的にsql.Dbへのポインタを渡す)
	// => handleRequestへ"依存性の源を注入"
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}

// 1. インタフェースTextを渡す
func handleRequest(t Text) http.HandlerFunc {
	// 2. 正しいシグネチャの関数を返す
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			// 3. 実際のハンドラにインタフェースTextを渡す
			err = handleGet(w, r, t)
		// case "POST":
		// 	err = handlePost(w, r, t)
		// case "PUT":
		// 	err = handlePut(w, r, t)
		// case "DELETE":
		// 	err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow(
		"select id, content, author from posts where id = $1",
		id,
	).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// 1. インタフェースTextを受け入れ
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 2. 構造体Postからデータを取得
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

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

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
