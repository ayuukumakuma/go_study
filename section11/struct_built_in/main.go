package main

import "fmt"

// 構造体
// 埋め込み

type T struct {
	User User // フィールド名と型が同名になることはgolangではよくある
	// User // フィールド名を省略することもできる
}

type User struct {
	Name string
	Age int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User.Name) // フィールド名を指定してアクセスする

	// line10のように省略した場合はフィールド名をそのまま使うことができる
	// fmt.Println(t.Name)

	t.User.SetName() // 内側のメソッドを呼び出すこともできる rubyと同じやね〜
	fmt.Println(t)
}
