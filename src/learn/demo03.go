package main

import (
	"fmt"
	"math"
)

/**
常量定义姿势（常量必须有 value）

	完全体：const name type = value
	后续的使用自动补类型：const name = value

枚举定义姿势（Go 没有枚举，用 const 块来定义）

	自定义枚举：const ( name1=value1 name2=value2 )
	iota 表达式枚举：const ( name1=iota表达式 name2 )

*/
// 常量
func consts() {
	// 指定类型
	const filename string = "filename-const"
	// 不指定类型，表示类型不定 （const数值可以用当各种类型来用）
	const a, b = 3, 4
	var c int
	// 由于类型不定，所以这里不需要强转，如果定义为 const a, b int = 3, 4，则需要强转
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c)
}

// 枚举
func enums() {
	// 使用 const 块来实现枚举
	const (
		java = 0
		cpp  = 1
		c    = 2
	)
	fmt.Println(java, cpp, c) // 0 1 2
	// 使用 iota 块来实现自增枚举
	const (
		java1 = iota
		cpp1
		c1
	)
	fmt.Println(java1, cpp1, c1) // 0 1 2
}

func main() {
	consts()
	enums()
}
