package main

import (
	"fmt"
)

/**
函数如果使用参数，该参数变量称为函数的形参。形参就像定义在函数体内的局部变量。调用函数，可以通过两种方式来传递参数。
即：值传递和引用传递，或者叫传值或传引用

（1）值传递：在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到原内容。
（2）引用传递：在调用函数时，将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到原内容数据。
引用传递的作用：#传指针使得多个函数能操作同一个对象
				#传指针更轻量（8Byte），只需要传内存地址。如果参数是非指针参数，那么值传递的过程中，每次在拷贝上面就会
             花费相对较多的系统开销（内存和时间）。所以当要传递打的结构体的时候，用指针是一个明智的选择。
				#Go语言中slice、map、chan类型的实现机制都是类似指针，所以可以直接传递，而不必取地址后传递指针
*/

/* 传数组类型*/
func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Printf("1、变量a的内存地址：%p,值为: %v \n", &a, a)
	changeArrayVal(a)
	fmt.Printf("2、调用changeArrayVal后，变量a的内存地址：%p,值为: %v \n", &a, a)
	//通过传指针，模仿传引用效果
	changeArrayPtr(&a)
	fmt.Printf("3、调用changeArrayPtr后，变量a的内存地址：%p,值为: %v \n", &a, a)
}

func changeArrayVal(a [4]int) {
	fmt.Printf("---------changeArrayVal函数内：值参数a的内存地址: %p,值为: %v \n", &a, a)
	a[0] = 99

}

func changeArrayPtr(a *[4]int) {
	fmt.Printf("---------changeArrayPtr函数内：指针参数a的内存地址: %p,值为: %v \n", &a, a)
	(*a)[1] = 233
}
