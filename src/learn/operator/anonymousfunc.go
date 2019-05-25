package main

import "fmt"

/**
Go支持匿名函数，即在需要使用函数的时候，再定义函数，匿名函数没有函数名，只有函数体，函数可以被作为一种类型被赋值给变量
匿名函数也往往以变量方式被传递
匿名函数经常被用于实现回调函数、闭包等。
*/

/**
定义匿名函数
*/
//1.在定义时调用匿名函数
func main() {
	func(data int) {
		fmt.Println("Hello", data)
	}(100)
	test1()
}

func s(data int) {
	fmt.Println("Hello", data)
}

//2.将匿名函数赋值给变量
func test1() {
	f := func(data string) {
		fmt.Println(data)
	}
	f("test")
}
