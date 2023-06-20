package main

import (
	"fmt"
	"time"
)

// channel
// チャンネルとgoroutin

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	// データを送信してくれるgoroutinがないとデッドロックが発生する
	// channelは複数のgoroutin間でデータを共有させるための仕組み
	ch1 := make(chan int)
	ch2 := make(chan int)

	go reciever(ch1) // 並行処理
	go reciever(ch2) // chにデータが入ったらデータを表示させる

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}
