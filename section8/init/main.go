package main

import "fmt"

// init

// init関数を定義すると必ず最初に走るようになる
// その名の通り初期化に使用される
func init() {
	fmt.Println("init")
}

// 複数定義もできる(実行順は上から順)
// まとめて定義するのが一般的
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
