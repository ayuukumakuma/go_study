package main

import "fmt"

// 式によるスイッチ

func main() {
	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// switch n := 2; n { // switch文内で変数を定義できる スコープはswitch文内
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }

	n := 6
	switch {
	case n > 0 && n < 4: // 条件式を記述する場合には変数定義を省略できる 列挙型と条件式は共存できない
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}
}
