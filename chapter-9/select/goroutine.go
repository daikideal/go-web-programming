package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	// ループのたびにチャネルa, bのどちらから受け取るかを決める
	// (選択した時点でどちらに値が入っているかによるが、両方入っているとランダム)
	for i := 0; i < 5; i++ {
		// チャネルからの取り出しを待つ
		time.Sleep(1 * time.Microsecond)

		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default: // 全てのチャネルが停止しているときに実行される
			fmt.Println("Default")
		}
	}
}
