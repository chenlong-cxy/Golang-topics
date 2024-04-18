package main

import "fmt"

func main() {
	// // 算术运算符
	// var num1 = +25
	// var num2 = -10
	// var num3 = num1 + num2
	// var num4 = num1 - num2
	// var num5 = num1 * num2
	// var num6 = num1 / num2
	// var num7 = num1 % num2
	// num1++
	// num2--
	// var s = "2021" + "dragon"

	// fmt.Printf("num1 = %d\n", num1) // 26
	// fmt.Printf("num2 = %d\n", num2) // -11
	// fmt.Printf("num3 = %d\n", num3) // 15
	// fmt.Printf("num4 = %d\n", num4) // 35
	// fmt.Printf("num5 = %d\n", num5) //-250
	// fmt.Printf("num6 = %d\n", num6) // -2
	// fmt.Printf("num7 = %d\n", num7) // 5
	// fmt.Printf("s = %s\n", s)       // 2021dragon

	// // 关系运算符
	// var num1, num2 = 10, 20
	// fmt.Println(num1 == num2) // false
	// fmt.Println(num1 != num2) // true
	// fmt.Println(num1 < num2)  // true
	// fmt.Println(num1 > num2)  // false
	// fmt.Println(num1 <= num2) // true
	// fmt.Println(num1 >= num2) // false

	// // 逻辑运算符
	// var num1, num2 = 10, 20
	// fmt.Println(num1 > 0 && num2 > 0)   // true
	// fmt.Println(num1 > 10 || num2 > 10) // true
	// fmt.Println(!(num1 > num2))         // true

	// // 赋值运算符
	// var num1, num2 = 10, 20
	// num1 = num2  // num1 = 20, num2 = 20
	// num1 += 10   // num1 = 30, num2 = 20
	// num2 -= 10   // num1 = 30, num2 = 10
	// num1 *= 2    // num1 = 60, num2 = 10
	// num2 /= 5    // num1 = 60, num2 = 2
	// num1 %= 8    // num1 = 4, num2 = 2
	// num1 <<= 1   // num1 = 8, num2 = 2
	// num2 >>= 1   // num1 = 8, num2 = 1
	// num1 &= num2 // num1 = 0, num2 = 1
	// num1 |= 4    // num1 = 4, num2 = 1
	// num2 ^= num1 // num1 = 4, num2 = 5
	// fmt.Printf("num1 = %d, num2 = %d\n", num1, num2) // num1 = 4, num2 = 5

	// // 位运算符
	// var num1, num2 = 12, -4
	// // 00...01100 --> 12的原码、反码、补码：

	// // 10...00100 --> -4的原码
	// // 11...11011 --> -4的反码
	// // 11...11100 --> -4的补码
	// fmt.Println(num1 & num2) // 00...01100 -> 12
	// fmt.Println(num1 | num2) // 11...11100 -> -4
	// fmt.Println(num1 ^ num2) // 11...10000 -> 11...01111 -> 10...10000 -> -16
	// fmt.Println(num1 << 2)   // 48
	// fmt.Println(num2 >> 2)   // 11...11111 -> -1

	// 其他运算符
	var num = 10
	fmt.Printf("num的地址是%p\n", &num) // num的地址是0xc00000e0b8
	var ptr = &num
	*ptr = 20
	fmt.Printf("num = %d\n", num) // 20
}
