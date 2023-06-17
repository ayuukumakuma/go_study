package main

import "fmt"

// 型によるスイッチ

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int) // 型アサーション(interface型に格納した変数の方を復元する)
	fmt.Println(i + 2)
	// fmt.Println(x + 2) // interface型なので計算できない

	f := x.(float64) // 3はint型なのでエラーを吐く
	fmt.Println(f)

	f, isFloat64 := x.(float64) // 返り値を二つ設定することでエラーを出さずに真偽値で返すことができる
	fmt.Println(f, isFloat64)

	if x == nil { // 型アサーション
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) { // switch文を使った型アサーションの方が推奨されている
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) { // 値を使いたい場合は変数に格納する
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string ")
	}
}
