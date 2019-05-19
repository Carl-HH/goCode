package main

import (
	"fmt"
	"strings"
)

/*
函数变量的使用步骤及意义
1.定义一个函数类型
2.实现定义的函数类型
3.作为参数调用

*/

type caseFunc func(str string) string

func main() {
	str := "abcdefghijklmn"
	fmt.Println(StringToCase(str, opreateFunc))
}

func opreateFunc(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func StringToCase(str string, caseFun caseFunc) string {
	return opreateFunc(str)
}
