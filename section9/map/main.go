package main

import "fmt"

// マップ

func main() {
	// pythonでいう辞書型みたいな
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) // 登録されていない場合は初期値が返ってくる(nullではない)

	_, ok := m4[4]
	if !ok {
		fmt.Println("error")
	}

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	delete(m4, 3) // 削除
	fmt.Println(m4)

	fmt.Println(len(m4))

}
