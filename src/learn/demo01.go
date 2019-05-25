package main

import "fmt"

/**
 	变量定义：
	 四种姿势
		完全体：var name type
		类型推断：var name = value
		最简体：name := value（仅用于函数内变量，包内变量不行）
		变量聚合定义：var( name1=value1 name2=value2 )
*/
// 定义包内变量
var (
	aa = 1
	bb = "tianya"
)

// 定义变量，只使用默认初值
func variableZeroValue() {
	var a int    // 0
	var s string // ""
	var b bool   // false
	fmt.Println(a, s, b)
}

// 定义变量，赋初值
func variableInitialValue() {
	var a, b int = 1, 2
	b = 3
	var s string = "haha"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 最简定义变量方式
func variableShorter() {
	a, b, c, d := 3, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb)
}
