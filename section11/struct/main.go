package main

import "fmt"

// 構造体

type User struct {
	// フィールド
	Name string
	Age int
	// X, Y int // 複数のフィールドを定義する場合はカンマで区切る
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) { // ポインタ型を受け取る
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User // 構造体のインスタンスを作成
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{} // 暗黙的に定義する
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 20} // 初期値を設定することもできる
	fmt.Println(user3)

	user4 := User{"user4", 40} // 暗黙的に代入することもできる(順番は構造体の順番と同じ)
	fmt.Println(user4)

	user5 := User{Name: "user5"} // 代入していないフィールドは初期値が入る
	fmt.Println(user5)

	user6 := new(User) // newは構造体のポインタ型を返す
	fmt.Println(user6) // &{ 0}のアドレスが返される

	user7 := &User{} // こっちでも同じ(こっちのが一般的)
	fmt.Println(user7) // 関数の引数として構造体のポインタ型を受け取ることが多い

	UpdateUser(user1)
	UpdateUser2(user7)
	fmt.Println(user1) // 値渡しのため値は変わらない {user1 10}
	fmt.Println(*user7) // ポインタ渡しのため値が変わる {A 1000}
}
