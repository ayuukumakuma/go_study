package main

import (
	"fmt"
	"time"
)

// channel
// close

func reciever(name string, ch chan int) {
	for {
		fmt.Println("Standby...")
		i, ok := <-ch
		if !ok { // チャネルが空かつcloseになったら
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	// チャネルはcloseという状態を持っている
	ch1 := make(chan int, 2)
	// close(ch1) // 明示的に閉じる
	// ch1 <- 1 // 送信しようとしてもエラーになる
	// fmt.Println(<-ch1) // 受信はcloseでもできる

	// i, ok := <-ch1     // 受信は二つの返り値に割り当てることができる
	// fmt.Println(i, ok) // okにはチャネルの開閉状態を表すbooleanが入る．今回は閉じているのでclose
	// 正確にはチャネル内のバッファーが全て空いている状態かつ閉じていることが条件(すでにデータが入っていればtrueになる)

	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(10 * time.Second)
}
