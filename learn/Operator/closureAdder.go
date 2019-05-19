package main

import "fmt"

func main() {
	res := Counter()
	fmt.Printf("%T \n", res)
	fmt.Println("res:", res)
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	fmt.Println("res:", res())

	res = Counter()
	fmt.Printf("%T \n", res)
	fmt.Println("res:", res)
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	fmt.Println("res:", res())
}

//闭包函数，实现计数器功能
//闭包函数，实际返回函数的地址，闭包函数实现步骤:
//1.定义闭包函数（名字、参数、返回值<返回值一定是一个函数>）
//2.定义一个局部变量
//3.内部的匿名函数要返回操作后的局部变量
//4.整个函数返回当前这个函数
func Counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	fmt.Println("Counter内部:", res)
	return res
}
