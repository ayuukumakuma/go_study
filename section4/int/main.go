package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100

	/*
		環境によってサイズが変わる
		int8 -128 ~ 127
		int16 -32768 ~ 32767
		int32 -2147483648 ~ 2147483647
		int64 -9223372036854775808 ~ 9223372036854775807
	*/

	var i2 int64 = 200 // intとint64は別物として扱われる
	// fmt.Println(i + i2)

	fmt.Println(i + 50)

	fmt.Printf("%T\n", i2) // 型を表示する %Tは型を表示するフォーマット

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2)) // 型が違うと計算できない
}
