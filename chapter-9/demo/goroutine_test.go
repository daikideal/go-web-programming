package main

import (
	"testing"
	"time"
)

// 1. 通常の実行
func TestPrint1(t *testing.T) {
	print1()
}

// 2. ゴルーチンとして実行
// ※遅延させないとゴルーチンが何かを出力する前にテストケースが終了する
func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}

// 1. 通常の実行
func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

// 2. ゴルーチンとして実行
func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

// 1. 通常の実行
func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

// 2. ゴルーチンとして実行
func BenchmarkGoPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}
