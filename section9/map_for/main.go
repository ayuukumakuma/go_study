package main

import "fmt"

// マップ
// for

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	// 同様にkey-valueが返り値にくる<k, v>, <k>, <_, v>
	for k, v := range m {
		fmt.Println(k, v)
	}
}
