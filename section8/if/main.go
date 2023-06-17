package main

import "fmt"

// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println(("I don't know"))
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// if文内で定義した変数はスコープはif文内のみ
	x := 0
	if x := 2; true {
		fmt.Println(x) // このxは2
	}
	fmt.Println(x) // このxは0
}
