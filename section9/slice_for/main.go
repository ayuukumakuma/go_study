package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// 配列型と同じ
	for i, v := range sl { // <_, v>, <i>, <i, v>と色々指定できる
		fmt.Println(i, v)
	}

	// 古典的forでやってみる
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
