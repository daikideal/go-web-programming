package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		// 1. チャネルに入れる
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		// 2. チャネルから取り出す
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

// 数字が送信側から送られると、受信側がそれを取り出さない限り次の数字に進めない
// = バッファなしチャネル
// 送信側から送られた数字を、受信側は送られた順に取り出す
// = バッファ付きチャネル
func main() {
	// バッファなし
	// c := make(chan int)
	// バッファ付き
	c := make(chan int, 3)

	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
