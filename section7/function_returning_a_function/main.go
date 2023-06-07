package main

import "fmt"

// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := ReturnFunc() // fに関数を代入する
	f()               // 関数となったfを実行
}
