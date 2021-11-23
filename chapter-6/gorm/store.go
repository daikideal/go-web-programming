package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(
		postgres.Open("host=localhost user=gwp password=password dbname=gwp sslmode=disable"),
	)
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) // 1. => {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}

	Db.Create(&post)  // 2. 投稿の作成
	fmt.Println(post) // 3. => {1 Hello World! Sau Sheong [] 2015-04-12 11:38:50.91815604 +0800 SGT}

	comment := Comment{Content: "いい投稿だね！", Author: "Joe"} // 4. コメントの追加

	Db.Model(&post).Association("Comments").Append(&comment) // 5. 投稿へのコメントの取得

	var readPost Post
	Db.Where("author = ?", "Sau Sheong").Last(&readPost)

	var comments []Comment
	Db.Model(&readPost).Association("Comments").Find(&comments)
	fmt.Println(comments[0]) // 6. => {1 いい投稿だね！ Joe 1 2015-04-13 11:38:50.920377 +0800 SGT}
}
