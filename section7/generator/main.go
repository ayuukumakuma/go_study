package main

import "fmt"

// ジェネレーター
// 何らかのルールに従って連続した値を返し続ける仕組みの事。
// クロージャーを使って実装する

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()  // integers()の返り値である関数をintsに代入
	fmt.Println(ints()) // integersが実行中である限り変数は保持される
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers() // 別の変数に関数を代入すると環境としては別になる
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}
