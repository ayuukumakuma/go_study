package main

import "fmt"

type User struct {
	Name string
	Age int
}

type Users []*User // ユーザー型のslice型の宣言Users型

//　一応このようにも書けるが望ましくない
// type Users struct {
// 	Users []*User
// }

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{} // Users型の変数を定義

	// ポインタ型のユーザーを入れる(User型はだめ)
	// users = append(users, user1)
	users = append(users, &user1, &user2, &user3, &user4) // Users型に格納

	for _, u := range users {
		fmt.Println(*u) // *をつけないとポインタ型が表示される(&{user1 10})
	}

	users2 := make([]*User, 0) // makeで初期化することもできる
	users2 = append(users2, &user1, &user2, &user3, &user4)

	for _, u := range users {
		fmt.Println(*u)
	}
}
