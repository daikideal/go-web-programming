// 1. テストファイルはテストされる関数と同じパッケージに置く
package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	// 2. テストされる関数の呼び出し
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	// 3. 結果が予想どおりチェックし、違えばエラーメッセージを設定
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	// 3. 結果が予想どおりチェックし、違えばエラーメッセージを設定
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	// 4. テストを全て省略
	t.Skip("Skipping encoding for now")
}

func TestLognRunningTest(t *testing.T) {
	// -short フラグでスキップすることを指定
	// https://pkg.go.dev/testing#Short
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
