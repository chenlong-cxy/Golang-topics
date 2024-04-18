package main

import "fmt"

func main() {
	// // 输出一行数据
	// fmt.Println("Hello", "Golang") // Hello Golang

	// // 格式化输出
	// var name string = "2021dragon"
	// fmt.Printf("your id is %s\n", name) // your id is 2021dragon

	// 输入一行数据
	var str1 string
	fmt.Scanln(&str1)

	// 格式化输入
	var str2 string
	fmt.Scanf("%s\n", &str2)

	fmt.Printf("%s\n", str1)
	fmt.Printf("%s\n", str2)
}
