package main

import "fmt"

// 多变量声明
var (
	s1 string = "dragon"
	n1 int    = 10
	s2        = "dragon"
	n2        = 10
)

func main() {
	// 声明变量
	// var s string       // 声明单个变量
	// var n1, n2, n3 int // 声明多个同类型变量

	// // 声明单个变量并赋值
	// var s1 string = "dragon"
	// var s2 = "dragon"
	// s3 := "dragon"

	// // 声明多个同类型变量并赋值
	// var n1, n2, n3 int = 10, 20, 30
	// var n4, n5, n6 = 10, 20, 30
	// n7, n8, n9 := 10, 20, 30

	// // 声明多个不同类型变量
	// var s1, n1 = 10, "dragon"
	// s2, n2 := 10, "dragon"
	// var (
	// 	a = 10
	// 	b = 20
	// )
	var bool = 10
	fmt.Println(bool)

}
