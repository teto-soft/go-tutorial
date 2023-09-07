package main

import "fmt"

func main() {
	x := 10
	y := 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// 関係演算子
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	// 論理演算子
	fmt.Println(x > 5 && x < 15)
	fmt.Println(x > 5 || x < 15)

	// 複合代入演算子
	z := 8
	z += 12
	fmt.Println(z)

	// インクリメント演算子
	a := 8
	b := 8
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)
}
