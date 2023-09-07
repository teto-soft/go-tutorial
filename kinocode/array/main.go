package main

import "fmt"

func main() {
	a := [3]string{"A", "B", "C"}
	a[1] = "D"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	// 配列の要素数を省略する
	b := [...]string{"sato", "suzuki", "takahashi"}
	b[1] = "tanaka"

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

	c := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}

	fmt.Println(c[0][0])
	fmt.Println(c[0][1])
	fmt.Println(c[1][0])
	fmt.Println(c[1][1])
}
