package main

import "fmt"

func main() {
	res := adder()
	fmt.Printf("%T \n", res)
	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(res(i))
	}
}

//实现计数器的闭包函数
func adder() func(int) int {
	sum := 0
	res := func(num int) int {
		sum += num
		return sum
	}
	return res
}
