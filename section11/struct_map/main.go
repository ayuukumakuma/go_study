package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	// mapの中に構造体を入れることもできる
	// valueにUser型を持ってくる
	m := map[int]User{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	fmt.Println(m)

	// keyにUser型を持ってくる
	m2 := map[User]string{
		{"user1", 10}: "Tokyo",
		{"user2", 20}: "Osaka",
	}

	fmt.Println(m2)

	// make関数を使うこともできる
	m3 := make(map[int]User) // この状態では空のmapが作られる
	m3[1] = User{"user1", 10}

	fmt.Println(m3)

	for _, v := range m { // mapの中身を表示する
		fmt.Println(v)
	}
}
