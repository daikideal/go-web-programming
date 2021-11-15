package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// ユニークなIDを使って投稿を取得(= Idを投稿へのポインタに対応づける)
var PostById map[int]*Post
// 著書名を使って投稿を取得(= 著書名を投稿へのポインタのスライスに対応づける)
var PostsByAuthor map[string][]*Post

// 注意: 投稿そのものへの対応づけではない。
// 			取得したいのは同一の投稿であって異なるコピーではないため

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
