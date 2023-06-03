package main

import "fmt"

// 型
// 配列型

func main() {
	var arr1 [3]int // 配列の定義
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) // 要素数までが型になるので変えることができない

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"} // 要素数を省略すると自動で要素数が決まる
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0]) // 要素の取得
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr2[2] = "C" // 要素の更新
	fmt.Println(arr2)

	// 要素数が異なる配列同士の代入はできない

	fmt.Println(len(arr1)) // 要素数の取得
}
