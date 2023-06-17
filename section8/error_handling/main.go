package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング

func main() {
	var s string = "100"
	// i, _ := strconv.Atoi(s) // _にエラーが入っている(この場合は変数未使用のエラーを回避している)

	i, err := strconv.Atoi(s) // エラーが発生した場合にはerrにエラーが格納されるため，if文でエラーを受け取れる
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i = %T\n", i)
}
