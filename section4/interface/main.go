package main

import "fmt"

// 型
// interface & nil

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	x = 2
	// fmt.Println(x + 3) // interface型は演算できない
}
