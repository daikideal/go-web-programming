package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	// 1. データを入れて構造体を作成する
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	// 2. 構造体を組み換えて(marshal)バイト列のXMLデータにする
	// output, err := xml.Marshal(&post)
	// if err != nil {
	// 	fmt.Println("Error marshalling to XML:", err)
	// 	return
	// }
	// err = ioutil.WriteFile("post.xml", output, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing XML to file:", err)
	// 	return
	// }

	// 2. 見栄えを良くし、XML宣言を付与する場合
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
