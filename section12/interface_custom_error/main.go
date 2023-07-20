package main

import "fmt"

// interface
// カスタムエラー

//　このようなerror型が定義されている
// type error interface {
// 	Error() string
// }

// このようにして独自のErrorを定義できる(golang本来のErrorと共通の性質を持つ)
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError() // errはerror型なのでMyErrorのMessageやErrCodeにはアクセスできない
	fmt.Println(err.Error())

	// アクセスするときはこうする
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
