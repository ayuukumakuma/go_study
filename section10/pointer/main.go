package main

import "fmt"

// ポインタ

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n) // アドレスを表示

	Double(n)
	fmt.Println(n) // 引数として渡したnの値は変わらない

	var p *int = &n // *int型のポインタを定義 p = &n
	fmt.Println(p) // nのアドレスが表示される
	fmt.Println(*p) // nの値が表示される 100

	/*
	*p = 300
	fmt.Println(n) // nの値が変わる 300
	n = 200
	fmt.Println(*p) // pの値が変わる 200
	*/

	Doublev2(&n) // 参照渡し(ポインタを渡す)をすることでコピーせずに値を変更できる
	fmt.Println(n) // 400
	Doublev2(p) // こっちでも同じ結果になる
	fmt.Println(*p) // 800

	var sl []int = []int{1, 2, 3} // 参照型はもともと参照渡しの機能を持っている
	Doublev3(sl)
	fmt.Println(sl) // [2, 4, 6]

}
