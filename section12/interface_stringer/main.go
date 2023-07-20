package main

import "fmt"

// interface
// fmt.Stringer

// このようなinterfaceが定義されている(stringを返すだけ)
// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

// メソッドの上書きみたいな感じ
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B) // <<100, ABC>>と表示される
}

func main() {
	p := &Point{100, "ABC"} // 上書きしないとこのまま&{100 ABC}と表示される
	fmt.Println(p)
}
