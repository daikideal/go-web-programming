package main

import (
	"testing"
)

// デコード関数のベンチマーク
func BenchmarkDecode(b *testing.B) {
	// 1. 関数をb.N回繰り返してベンチマークを得る
	// コードが実行されるとb.Nは必要に応じて変化する(ユーザーは直接回数を指定できない、実行制限時間は指定可能)
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
