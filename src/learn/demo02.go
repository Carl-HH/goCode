package main

import (
	"fmt"
	"math"
)

/**
内建变量类型
	8 类（指针单独划了出来）
	bool
	string
	(u)int、(u)int8、(u)int16、(u)int32、(u)int64 （无符号和有符号整数）
	uintptr 指针
	byte
	rune 字符型，32 位，类比 char
	float32、float64
	complex32、complex64 复数 i = √-1
*/

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	// float64 和 int 可以不加小括号，也可以加上
	// 开方内建函数定义：func Sqrt(x float64) float64
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	triangle()
}
