package main

import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
	test
	`) // 複数行の文字列

	fmt.Println("\"") // エスケープシーケンス
	fmt.Println(`"`)  // エスケープシーケンスを無視

	fmt.Println(s[0])         // 文字列の1文字目を取得 byte型で返ってくる
	fmt.Println(string(s[0])) // 文字列の1文字目を取得 string型に変換して返す
}
