package main

import "fmt"

// panic recover

func main() {
	defer func() {
		// recoverにはpanicが発生すると値が格納される
		// 基本的にdeferと合わせて使う(エラーを発生させているためdeferでないとキャッチできない)
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	// 例外処理(実際はエラーハンドリングを使った方がいいと推奨されている)
	panic("runtime error") // 強制的にエラーを発生させる
}
