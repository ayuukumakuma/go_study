package main

import "fmt"

// スライス
// copy

func main() {
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl) // 上のように渡すとslとsl2で同じメモリのアドレスが参照されてしまい，同期された状態になる(参照型特有)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	n := copy(sl2, sl) // nにはコピーに成功した要素数が格納される(今回なら5
	fmt.Println(n, sl2)
}
