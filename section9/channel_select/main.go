package main

import "fmt"

// channel
// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1) // 空のチャネルから受信することはできない
	// fmt.Println(v2) // ↑がデッドロックになるとここも実行されなくなってしまう

	// selectは複数のチャネルの送受信をgoroutinを停止させることなく制御できる

	select {
	// caseにはチャネルに関する処理しか書けない(エラーを吐いてしまう)
	case v1 := <-ch1: // ch1にデータが入った場合はこっち
		fmt.Println(v1 + 1000)
	case v2 := <-ch2: // ch2にデータが入った場合はこっち
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("I don't know")
	}

	// 活用例
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
	// golangは初めから非同期処理を効率的に処理できるようにデザインされている
}
