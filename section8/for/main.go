package main

import "fmt"

// for

func main() {
	// for { // 条件を指定しないと無限ループ
	// 	fmt.Println("Loop")
	// }

	// i := 0
	// for { // 条件を内部で指定してbreakする方法
	// 	i++
	// 	if i == 20 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 { // while文みたいな使い方
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ { // 古典的forって呼ぶらしい これが一番使うかも
	// 	if i == 3 {
	// 		continue // continueでその後の処理を飛ばし，次のループに移る
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ { // 配列をfor文で回す
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1, 2, 3}
	// for i, v := range arr { // iにはインデックスが，vには配列の要素が格納される
	// 	fmt.Println(i, v)
	// }

	// arr := [3]int{1, 2, 3}
	// for _, v := range arr { // インデックスが必要ない場合はアンスコにする
	// 	fmt.Println(v)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := range arr { // インデックスのみでいい場合はiのみにする
	// 	fmt.Println(i)
	// }

	// 簡単にスライスを触る
	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// 簡単にマップを触る(pythonの辞書みたいな)
	m := map[string]int{"apple": 100, "banana": 200} // map[<key>]<value>
	for k, v := range m {
		fmt.Println(k, v)
	}
}
