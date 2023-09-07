package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		if i == 0 {
			continue
		}
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// 配列の要素を順番に取り出す
	a := [...]string{"sato", "suzuki", "takahashi"}
	for j := 0; j < len(a); j++ {
		fmt.Println(j, a[j])
	}

	for i, v := range a {
		fmt.Println(i)
	}
}
