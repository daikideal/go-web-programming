package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>content"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// XMLを要素毎にデコード
	decoder := xml.NewDecoder(xmlFile) // 1. XMLデータからdecoderを生成
	for {                              // 2. decoder内のXMLデータを順次処理
		t, err := decoder.Token() // 3. 各処理でdecoderからToken(XML要素を表すインタフェース)を取得
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) { // 4. Tokenの型をチェック
		case xml.StartElement: // StartElement = XML要素の開始タグ
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se) // 5. XMLデータを構造体にデコード
			}
		}
	}
}
