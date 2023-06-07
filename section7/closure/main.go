package main

import "fmt"

// クロージャー
// 関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// Laterが実行されている間はずっと値は保持されている
func main() {
	f := Later() // Later()の返り値は引数と返り値がstringの無名関数．これをfに代入している
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))

}
