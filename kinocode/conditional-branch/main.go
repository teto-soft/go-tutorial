package main

import "fmt"

func main() {
	x := 10
	y := 12

	// if文の中で変数を宣言することができる
	// 宣言された変数は if 文の中でしか使えない
	if age := x + y; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
