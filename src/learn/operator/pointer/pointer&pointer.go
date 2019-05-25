package main

import "fmt"

/**
指针的指针
1.如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量的时候，第一个指针存放第二个指针的地址，第二个指针存放变量的地址
格式：var ptr **int
*/
func main() {
	var a int = 1234
	var ptr *int = &a
	var pptr **int = &ptr

	fmt.Println(a, ptr, pptr)

	fmt.Println(*ptr, *pptr, **pptr)

}
