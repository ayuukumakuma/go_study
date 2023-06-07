package main

import "fmt"

// 関数

func Plus(x int, y int) int { // 戻り値の型も指定する
	return x + y
}

// func Plus(x, y int) int { // 引数が同じ型の時は型を省略できる
// 	return x + y
// }

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4) // 複数ある戻り値(変数)をいずれかだけ使いたい場合は_を使って不使用のエラーを避ける
	fmt.Println(i2)

	i4 := Double(1000) // 返り値の変数を指定できる. 明示的に指定できる
	fmt.Println(i4)

	Noreturn() // 引数も返り値もなし
}
