package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var num1 int = 10
	// //var num2 float64 = num1 // error:类型不同，无法赋值
	// var num2 float64 = 1.2
	// // var ret = num1 * num2 // error:类型不同，无法参与运算
	// fmt.Printf("num1 = %d, num2 = %.2f\n", num1, num2) // num1 = 10, num2 = 1.20

	// // 基本数据类型相互转换
	// var num int = 10
	// var num1 float64 = float64(num)
	// var num2 int8 = int8(num)
	// var num3 uint16 = uint16(num)
	// fmt.Printf("num1 = %.2f, num2 = %d, num3 = %d\n", num1, num2, num3)

	// // 基本数据类型转string
	// var i1 int = 10
	// var i2 float32 = 3.14
	// var i3 bool = true
	// var i4 uint = 20
	// var s string = fmt.Sprintf("%d %.2f %t %d", i1, i2, i3, i4)
	// fmt.Printf("s = %q\n", s) // s = "10 3.14 true 20"

	// // 基本数据类型转string
	// var i1 int = 10
	// var i2 float32 = 3.14
	// var i3 bool = true
	// var i4 uint = 20
	// var s1 string = strconv.FormatInt(int64(i1), 10)
	// var s2 string = strconv.FormatFloat(float64(i2), 'f', -1, 32)
	// var s3 string = strconv.FormatBool(i3)
	// var s4 string = strconv.FormatUint(uint64(i4), 10)
	// fmt.Printf("s1 = %q\n", s1) // s1 = "10"
	// fmt.Printf("s2 = %q\n", s2) // s2 = "3.14"
	// fmt.Printf("s3 = %q\n", s3) // s3 = "true"
	// fmt.Printf("s4 = %q\n", s4) // s4 = "20"

	// // string转基本数据类型
	// var s1 string = "10"
	// var s2 string = "3.14"
	// var s3 string = "true"
	// var s4 string = "20"
	// i1, err1 := strconv.ParseInt(s1, 10, 0)
	// i2, err2 := strconv.ParseFloat(s2, 32)
	// i3, err3 := strconv.ParseBool(s3)
	// i4, err4 := strconv.ParseUint(s4, 10, 0)
	// fmt.Printf("i1: value = %d, type = %T, err = %v\n", i1, i1, err1)   //i1: value = 10, type = int64, err = <nil>
	// fmt.Printf("i2: value = %.2f, type = %T, err = %v\n", i2, i2, err2) //i2: value = 3.14, type = float64, err = <nil>
	// fmt.Printf("i3: value = %t, type = %T, err = %v\n", i3, i3, err3)   // i3: value = true, type = bool, err = <nil>
	// fmt.Printf("i4: value = %d, type = %T, err = %v\n", i4, i4, err4)   // i4: value = 20, type = uint64, err = <nil>

	// // int转string
	// var num1 int = 10
	// s1 := strconv.Itoa(num1)
	// fmt.Printf("s1 value = %q, s1 type = %T\n", s1, s1)  // s1 value = "10", s1 type = string

	// // string转int
	// var s2 string = "20"
	// num2, err := strconv.Atoi(s2)
	// fmt.Printf("num2 value = %d, num2 type = %T, err = %v\n", num2, num2, err) // num2 value = 20, num2 type = int, err = <nil>

	// // 指针类型
	// var num int = 10
	// var ptr *int = &num
	// *ptr = 20
	// fmt.Printf("num = %d\n", num)                     // num = 20
	// fmt.Printf("address = %p\n", ptr)                 // address = 0xc00000e0b8
	// fmt.Printf("ptr size = %d\n", unsafe.Sizeof(ptr)) // ptr size = 8

	var arr [6]int
	fmt.Printf("%d\n", unsafe.Sizeof(arr))
}
