package main

import "fmt"

/**
当一个指针被定义后没有分配到任何变量的时候，它的值为nil
nil也叫空指针
nil在概念上和其他语言的null，NONE，NULL一样，都指代零值或空值
一个指针变量通常缩写为ptr
*/
func main() {
	var ptr *int
	fmt.Printf("ptr的类型为%T，ptr的值为%v \n", ptr, ptr)
}
