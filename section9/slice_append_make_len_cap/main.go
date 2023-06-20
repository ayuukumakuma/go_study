package main

import "fmt"

// スライス
// append make len cap

func main() {
	// スライスは可変長の要素数
	sl := []int{100, 200}
	fmt.Println(sl)

	// append
	sl = append(sl, 300, 400, 500)
	fmt.Println(sl)

	// make
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// len
	fmt.Println(len(sl2))

	// cap(容量)
	// メモリを気にするような開発で使う
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
}
