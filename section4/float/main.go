package main

import "fmt"

// 型
// 浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // 環境依存ではなく，常にfloat64
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2 // あまり使わない
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32)) // 型変換

	// 0での除算でもエラーにならない
	zero := 0.0
	pinf := 1.0 / zero // 非数
	fmt.Println(pinf)

	ninf := -1.0 / zero // 非数
	fmt.Println(ninf)

	nan := zero / zero // 非数
	fmt.Println(nan)

	// 他にも自然数(uint), 複素数(complex)などがある
	// var u8 uint = 255 //uint型

	// var c64 complex64 = -5 + 12i //complex型
}
