package main

import "fmt"

// channel
// for

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) // チャネルでfor文を使うときはcloseしてから回すのが基本

	for i := range ch1 { // 空のチャンネルからデータを受け取るとデッドロックが発生する
		fmt.Println(i)
	}
}
