package main

import (
	"fmt"
	"os"
)

// defer

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func Rundefer() {
	// 後から評価されるので3, 2, 1と表示される
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	// 関数の終了時に実行できる処理を登録できる
	TestDefer()
	// deferで定義したものが関数の最後に実行される
	// START
	// END

	// 複数行の場合は無名関数
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	Rundefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// リソースの解放処理などに使う
	defer file.Close()

	file.Write([]byte("Hello"))
}
