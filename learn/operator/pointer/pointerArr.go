package main

import "fmt"

/**
指针数组：元素为指针的数组
定义一个指针数组，如：var ptr [3]*string
有一个元素个数相同的数组，将该数组中的每个元素的地址赋值给该指针数组。也就是说该指针数组与某一个数组完全对应。
可以通过*指针变量获取到该地址所对应的数值
*/

const COUNT int = 4

func main() {
	a := [COUNT]string{"a", "b", "c", "d"}
	//查看数组的指针的类型和值
	fmt.Printf("%T , %v \n", &a, &a)
	//定义指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T , %v \n", ptr, ptr)
	for i := 0; i < COUNT; i++ {
		ptr[i] = &a[i]
	}
	fmt.Printf("%T , %v \n", ptr, ptr)
	//fmt.Printf("%T , %v \n", ptr, ptr)
	//根据指针数组元素的每个地址，获取该地址所指向的元素的真实数值
	for i := 0; i < COUNT; i++ {
		fmt.Printf("%T , %v \n", *ptr[i], *ptr[i])
	}

}
