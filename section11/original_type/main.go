package main

import "fmt"

type MyInt int // 独自型の宣言

// 独自型にメソッドを生やすこともできる
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi) // 0
	fmt.Printf("%T\n", mi) // main.MyInt

	// MyIntとintは別物なので計算できない
	// i := 100
	// fmt.Println(mi + i) // コンパイルエラー

	mi.Print() // 0
}
