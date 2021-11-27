package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	// 1. データの入った構造体を作成する
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	// 2. データを保存するXMLファイルを作成する
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	// 3. XMLファイルに対してエンコーダを生成する
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	// 4.構造体をファイルにエンコードする
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
