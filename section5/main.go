package main

import "fmt"

// 定数

const Pi = 3.14 // 頭文字を大文字にすると他のパッケージからも参照できる

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	// 値を指定しないと省略できる
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1 // 定数は型をオーバーフローしても定義できるが関数内で使うときには型の範囲内に従わなければならない

const (
	c0 = iota // iotaは連番を表す
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// 再定義不可
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)
	// fmt.Println(Big) // オーバーフロー

	fmt.Println(c0, c1, c2)
}
