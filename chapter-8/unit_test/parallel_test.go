package main

import (
	"testing"
	"time"
)

// 1. 実行時間1秒の作業
func TestParallel_1(t *testing.T) {
	// 2. テストケースを並行実行するため関数Parallelの呼び出し
	t.Parallel()
	time.Sleep(1 * time.Second)
}

// 3. 実行時間2秒の作業
func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

// 4. 実行時間3秒の作業
func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
