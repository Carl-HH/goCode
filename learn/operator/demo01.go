package main

import "fmt"

/*
  变量声明及初始化
*/
func main() {
	var (
		//标准格式
		a int
		b string
		c float32
		d []int
		e int32 = 100
		f [3]string
		g bool
		h func() bool
		//自动推断类型格式
		l = 200
	)
	//短变量声明格式（首选编译器自动推断，但只能用在函数体中，不可用于全局变量的声明和赋值，必须是没有定义过的变量）（在多个短声明
	// 和赋值中，至少有一个新的声明出现在左侧中，那么即便有其他变量名可能是重复声明，编译器也不会报错 eg：m,n,q := "a","b",10）
	//格式： 变量名:= 表达式
	m := "19"
	n := "190"
	m, n, q := "a", "b", 10
	//匿名变量（"_"代替）

	//格式化输出
	fmt.Printf("%T , %v \n", a, a)
	fmt.Printf("%T , %q \n", b, b)
	fmt.Printf("%T , %v \n", c, c)
	fmt.Printf("%T , %v \n", d, d)
	fmt.Printf("%T , %v \n", e, e)
	fmt.Printf("%T , %q \n", f, f)
	fmt.Printf("%T , %v \n", g, g)
	fmt.Printf("%T , %v \n", h, h)
	fmt.Printf("%T , %v \n", l, l)
	fmt.Printf("%T , %v \n", m, m)
	fmt.Printf("%T , %v \n", n, n)
	fmt.Printf("%T , %v \n", q, q)

	//多重赋值（实现变量交换）
	x := 10
	y := 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}
