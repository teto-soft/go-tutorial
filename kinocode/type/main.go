package main

import (
	"fmt"
	"reflect"
)

func main() {
	num01:=123
	var num02 int = 1234567890
	num03:=1.23
	var num04 float64 = 1.234567890
	var str string = "Hello World"
	var flag bool = true

	// reflect.TypeOf()で型を取得できる
	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(flag))
}