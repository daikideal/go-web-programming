package main

import "fmt"

func callerA(c chan string) {
	c <- "Hello World!"
	// 1. 関数が呼び出されたらチャネルを閉じる
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	// 1. 関数が呼び出されたらチャネルを閉じる
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		// 2. チャネルが閉じているとok1とok2はfalseになる
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
