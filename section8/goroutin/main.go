package main

import (
	"fmt"
	"time"
)

// goroutin
// 並行処理

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// goroutinの構文を使うことで暗黙的かつ簡単に並行処理を行なってくれる
	go sub() // goをつけるだけ

	for {
		fmt.Println("Main Loop")
		time.Sleep(2000 * time.Millisecond)
	}
}
