package main

import "fmt"

// 構造体
// メソッド

type User struct {
	Name string
	Age int
}

func (u User) SayName() { // レシーバーを定義する rubyと同じやね〜
	fmt.Println(u.Name)
}

// 基本的にレシーバーはポインタ型を使うのが望ましい
func (u *User) SetName(name string) { // ポインタ型をレシーバーにすることで値を変更できる
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName() // メソッドを呼び出す

	user1.SetName("A")
	user1.SayName() // 宣言側でポインタ型か値型かが決まる

	user2 := &User{Name: "user2"} // ポインタ型で宣言しても同じ
	user2.SetName("B")
	user2.SayName()
}
