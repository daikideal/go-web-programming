package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	// "strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

// 事前処理
func setUp() {
	// // 1. テストを実行するマルチプレクサを生成
	// mux = http.NewServeMux()
	// // 2. テスト対象のハンドラを付加
	// mux.HandleFunc("/post/", handleRequest)
	// // 3. 返されたHTTPレスポンスを取得
	// writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	// 1. Postの代わりにFakePostを渡す
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

// func TestHandlePut(t *testing.T) {
// 	// JSONのコンテンツを送らなければならない
// 	json := strings.NewReader(`{"content":"Updated post", "author":"Sau Sheong"}`)
// 	request, _ := http.NewRequest("PUT", "/post/1", json)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// }
