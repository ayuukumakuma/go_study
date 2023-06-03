package main

import "fmt"

// 型
// byte(uint8の別名)型

func main() {
	byteA := []byte{72, 73} // byte型のスライス
	fmt.Println(byteA)

	fmt.Println(string(byteA)) // 文字列に変換
	//  文字列はアスキーコードで表現されている

	c := []byte("HI") // 文字列をbyte型に変換
	fmt.Println(c)
	fmt.Println(string(c))
}
