package main

import (
	"fmt"
	// . "fmt"とするとパッケージ名を省略できる (fmt.Println() -> Println())
	"go_udemy/section13/public_private_division/foo" // 自作fooパッケージのインポート
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version = "1.0"
	return AppName + " " + Version
}

// 引数と同じ名前の関数内の変数は再定義できない
// func Do(s string) (b string) {
// 	var b string = s
// 	return b
// }

func main() {
	fmt.Println(foo.Max) // 100
	// fmt.Println(foo.min) // 頭文字が小文字なのでエラーになる

	fmt.Println(foo.ReturnMin()) // パブリックな関数を通してプライベートのminを呼び出す

	fmt.Println(appName())
	// fmt.Println(AppName) // スコープ外のため参照できない

}
