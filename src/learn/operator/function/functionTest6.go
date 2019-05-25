package main

import (
	"fmt"
)

/**
值传递和引用传递的注意细节
1.Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
    #拷贝的内容有时候是非引用类型，(int、string、bool、数组、struct属于非引用类型)，这样就在函数中无法修改原内容数据
	#有的是引用类型(指针、Struct、map、chan属于引用类型)，这样就可以修改原内容数据
2.是否可以修改原内容数据，和传值、传引用没有必然的关系，在C++中，传引用肯定是可以修改原内容数据的，在Go语言中，虽然只有
传值，但是我们也可以修改原内容数据，因为参数可以是引用类型
3.传引用和引用类型是两个概念。虽然Go语言只有传值一种方式，但是可以通过传引用类型变量达到和传引用一样的效果。
*/

/* 传切片类型 （注意切片与数组的区别）*/
type student struct {
	name string
	age  int8
}

func main() {
	a := student{"carl", 34}
	fmt.Printf("1、变量a的内存地址：%p,值为: %v \n", &a, a)
	changeStructVal(a)
	fmt.Printf("2、调用changeStructVal后，变量a的内存地址：%p,值为: %v \n", &a, a)
	//通过传指针，模仿传引用效果
	changeStructPtr(&a)
	fmt.Printf("3、调用changeStructPtr后，变量a的内存地址：%p,值为: %v \n", &a, a)
}

func changeStructVal(a student) {
	fmt.Printf("---------changeStructVal函数内：值参数a的内存地址: %p,值为: %v \n", &a, a)
	a.name = "Jonh"
}

func changeStructPtr(a *student) {
	fmt.Printf("---------changeStructPtr函数内：指针参数a的内存地址: %p,值为: %v \n", &a, a)
	(*a).age = 40
}
