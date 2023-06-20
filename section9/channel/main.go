package main

import "fmt"

// channel
// 複数のgoroutin間でのデータ受け渡しのためのデータ構造
// 宣言，操作

func main() {
	// 双方向
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)

	ch2 := make(chan int) // 暗黙的

	fmt.Println(cap(ch1)) // バッファサイズを調べる
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5) // バッファサイズを指定
	fmt.Println(cap(ch3))

	ch3 <- 1 // 送信
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	i := <-ch3 // 受信
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	// channelはqueueの性質を持っている(行列みたいな感じ)
	// 先入先出

	// バッファサイズを超えるとデッドロックが発生するので適切な値を設定してあげなければならない
	ch3 <- 1
	fmt.Println(<-ch3)
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
}
