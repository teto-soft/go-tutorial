package main

import "fmt"

func main() {
	// 変数に使える文字は、英数字とアンダースコアのみ
	// 変数名の先頭に数字は使えない
	// var num = 1
	num := 1
	num01, num_01 := 2, 3

	fmt.Println(num)
	fmt.Println(num01)
	fmt.Println(num_01)

	// 大文字と小文字は区別される
	// 1文字目が小文字の場合は、その変数はそのパッケージ内でのみ使用可能
	// 1文字目が大文字の場合は、その変数は他のパッケージからも使用可能
	// Println の1文字目が大文字なので、 fmt ではない mainパッケージからも使用可能
	NUM := 4
	fmt.Println(NUM)
}