package main

import (
	"fmt"
)

func main() {
	// // 整数类型
	// var num int16 = 1024
	// fmt.Printf("num value = %d\n", num)               // num value = 1024
	// fmt.Printf("num type = %T\n", num)                // num type = int16
	// fmt.Printf("num size = %d\n", unsafe.Sizeof(num)) // num size = 2

	// // 字符类型
	// var ch1 byte = 'a'
	// fmt.Printf("ch1 value = %c\n", ch1)               // ch1 value = a
	// fmt.Printf("ch1 type = %T\n", ch1)                // ch1 type = uint8
	// fmt.Printf("ch1 size = %d\n", unsafe.Sizeof(ch1)) // ch1 size = 1

	// var ch2 rune = '龙'
	// fmt.Printf("ch2 value = %c\n", ch2)               // ch2 value = 龙
	// fmt.Printf("ch2 type = %T\n", ch2)                // ch2 type = int32
	// fmt.Printf("ch2 size = %d\n", unsafe.Sizeof(ch2)) // ch2 size = 4

	// // 浮点数类型
	// var num float32 = 3.1415926
	// fmt.Printf("num value = %.2f\n", num)             // num value = 3.14
	// fmt.Printf("num type = %T\n", num)                // num type = float32
	// fmt.Printf("num size = %d\n", unsafe.Sizeof(num)) // num size = 4

	// // 十进制表示形式
	// var num1 = 3.1415926
	// var num2 = .1415926
	// fmt.Printf("num1 = %.2f\n", num1) // num1 = 3.14
	// fmt.Printf("num2 = %.2f\n", num2) // num2 = 0.14

	// // 科学计数法表示形式
	// var num3 = 3.1415926e2
	// var num4 = 3.1415926e-2
	// fmt.Printf("num3 = %f\n", num3) // num3 = 314.159260
	// fmt.Printf("num4 = %f\n", num4) // num4 = 0.031416

	// // 复数类型
	// // var z complex64 = 1 + 2i
	// var z complex64 = complex(1, 2)
	// fmt.Printf("z value = %v\n", z)               // z = (1+2i)
	// fmt.Printf("z type = %T\n", z)                // z type = complex64
	// fmt.Printf("z size = %d\n", unsafe.Sizeof(z)) // z size = 8
	// fmt.Printf("z = %f\n", real(z))               // z的实部 = 1.000000
	// fmt.Printf("z = %f\n", imag(z))               // z的虚部 = 2.000000

	// // 布尔类型
	// var flag bool = true
	// fmt.Printf("flag value = %t\n", flag)               // flag value = true
	// fmt.Printf("flag type = %T\n", flag)                // flag type = bool
	// fmt.Printf("flag size = %d\n", unsafe.Sizeof(flag)) // flag size = 1

	// // string类型
	// var s string = "Hello World"
	// fmt.Printf("s value = %s\n", s)               // s value = Hello World
	// fmt.Printf("s type = %T\n", s)                // s type = string
	// fmt.Printf("s size = %d\n", unsafe.Sizeof(s)) // s size = 16

	// var s string = "Hello World"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("s[%d]: %c\n", i, s[i]) // 访问string的第i个元素
	// }

	// var s string = "Hello 世界"
	// tmp := []rune(s)
	// for i := 0; i < len(tmp); i++ {
	// 	fmt.Printf("tmp[%d]: %c\n", i, tmp[i]) // 访问rune的第i个元素
	// }

	// var s string = "Hello 世界"
	// tmp := []rune(s)
	// for i, v := range tmp {
	// 	fmt.Printf("tmp[%d]: %c\n", i, v) // 访问tmp的第i个元素
	// }

	// var s string = "Hello 世界"
	// for i, v := range s {
	// 	fmt.Printf("s[%d]: %c\n", i, v) // 访问s的第i个元素
	// }

	// var s1 string = "hello\nworld" // 双引号的方式表示字符串
	// fmt.Printf("s1 = %s\n", s1)
	// var s2 string = `hello\nworld` // 反引号的方式表示字符串
	// fmt.Printf("s2 = %s\n", s2)

	// // 字符串拼接
	// var s string = "Hello" + "World"
	// fmt.Printf("s = %s\n", s) // s = HelloWorld

	// const s1 string = "Hello"               // 声明常量
	// const s2 = "World"                      // 声明常量时省略类型
	// const s3, s4 string = "Hello", "Golang" // 声明多个同类型常量
	// const year, name = 2021, "dragon"       // 声明多个不同类型常量
	// fmt.Println(s1, s2)                     // Hello World
	// fmt.Println(s3, s4)                     // Hello Golang
	// fmt.Println(year, name)                 // 2021 dragon

	// const num1 = 10
	// const num2 = 3.14
	// const num3 = 1 + 2i
	// const ch = 'a'
	// const flag = true
	// const s = "dragon"
	// fmt.Printf("num1 type = %T\n", num1) // num1 type = int
	// fmt.Printf("num2 type = %T\n", num2) // num2 type = float64
	// fmt.Printf("num3 type = %T\n", num3) // num3 type = complex128
	// fmt.Printf("ch type = %T\n", ch)     // ch type = int32
	// fmt.Printf("flag type = %T\n", flag) // flag type = bool
	// fmt.Printf("s type = %T\n", s)       // s type = string

	// const (
	// 	a = 1
	// 	b
	// 	c = 2
	// 	d
	// )
	// fmt.Printf("a = %d\n", a) // a = 1
	// fmt.Printf("b = %d\n", b) // b = 1
	// fmt.Printf("c = %d\n", c) // c = 2
	// fmt.Printf("d = %d\n", d) // d = 2

	// const (
	// 	a = iota
	// 	b
	// 	c
	// 	d
	// )
	// fmt.Printf("a = %d\n", a) // a = 0
	// fmt.Printf("b = %d\n", b) // b = 1
	// fmt.Printf("c = %d\n", c) // c = 2
	// fmt.Printf("d = %d\n", d) // d = 3

	const (
		a = 1 << iota
		b
		c
		d
	)
	fmt.Printf("a = %d\n", a) // a = 1
	fmt.Printf("b = %d\n", b) // b = 2
	fmt.Printf("c = %d\n", c) // c = 4
	fmt.Printf("d = %d\n", d) // d = 8
}
