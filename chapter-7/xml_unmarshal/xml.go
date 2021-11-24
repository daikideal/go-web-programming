package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// 注意: 構造体タグの値のカンマ後にスペースを入れると意図しない挙動になる

type Post struct { // #A
	XMLName xml.Name `xml:"post"` // XML要素名を保存
	Id      string   `xml:"id,attr"` // XML要素の属性を保存
	Content string   `xml:"content"`
	Author  Author   `xml:"author"` // 下位要素を対応付ける
	Xml     string   `xml:",innerxml"` // RawXMLを得る
}

type Author struct { // 1. データを表す構造体を定義する
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` // XML要素の文字データを保存
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// XMLファイル全体を文字列として一気に解析
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post) // 2. XMLデータを構造体に格納する
	fmt.Println(post)
}
