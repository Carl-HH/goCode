package main

import "fmt"

/**
 *  1.指针是存储另一个变量的内存地址的变量，变量是一种使用方便的占位符，变量都指向计算机的内存地址，
一个指针变量可以用指向任何一个值的内存地址
    2.获取变量地址 取地址符 & 返回该变量的内存地址
	3.Go语言指针的特点，指针不能运算
	4.指针使用流程：定义指针变量，为指针变量赋值，访问指针变量中指向地址的值，获取指针的值：在指针类型的变量前加上“*”
来获取指针所指向的内容。获取一个指针意味着访问指针指向的变量的值。语法是：*a
*/
func pointerTest() {
	var a int = 100
	var ip *int = &a
	fmt.Printf("a的类型是%T, 值是%v \n", a, a)
	fmt.Printf("&a的类型是%T, 值是%v \n", &a, &a)
	fmt.Printf("ip的类型是%T, 值是%v \n", ip, ip)
	fmt.Printf("*ip的类型是%T, 值是%v \n", *ip, *ip)
	fmt.Printf("*&a的类型是%T, 值是%v \n", *&a, *&a)
	fmt.Println(a, &a, *&a)
	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
}

type Student struct {
	name string
	age  int
	sex  int8
}

func main() {
	pointerTest()

	typeStructTest()

}

func typeStructTest() {
	s1 := Student{"a", 30, 1}
	s2 := Student{"b", 25, 0}
	var pa *Student = &s1
	var pb *Student = &s2
	fmt.Println("\n---------------------")
	fmt.Printf("s1的类型为%T,值为%v \n", s1, s1)
	fmt.Printf("s2的类型为%T,值为%v \n", s2, s2)

	fmt.Printf("pa的类型为%T,值为%v \n", pa, pa)
	fmt.Printf("pb的类型为%T,值为%v \n", pb, pb)

	fmt.Printf("*pa的类型为%T,值为%v \n", *pa, *pa)
	fmt.Printf("*pb的类型为%T,值为%v \n", *pb, *pb)

	fmt.Println(pa.age, pa.name, pa.sex)
	fmt.Println((*pa).age, (*pa).name, (*pa).sex)
	fmt.Println(&pa.age, &pa.name, &pa.sex)
}
