package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	// 1. カウンタを減算
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	// 2. カウンタを減算
	wg.Done()
}

func main() {
	// 3. WatiGroupの宣言
	var wg sync.WaitGroup
	// 4. カウンタの初期化
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	// 5. カウンタが0になるまで実行を中断(ブロック)
	wg.Wait()
}
