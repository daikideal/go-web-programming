package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

// 1. Dbに接続
func init() {
	var err error
	Db, err = sql.Open(
		"postgres",
		"user=gwp dbname=gwp password=password sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
}

// 2. 投稿を1つだけ取り出す
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow(
		"select id, content, author from posts where id = $1",
		id,
	).Scan(
		&post.Id,
		&post.Content,
		&post.Author,
	)
	return
}

// 3. 新しい投稿の作成
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// 4. 投稿の更新
func (post *Post) update() (err error) {
	_, err = Db.Exec(
		"update posts set content = $2, author = $3 where id = $1",
		post.Id,
		post.Content,
		post.Author,
	)
	return
}

// 5. 投稿の削除
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
