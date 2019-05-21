package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	fmt.Println(a)
	fmt.Println("-----------------")
	modify(&a)
	fmt.Println(a)
	modify2(a)
	fmt.Println(a)
}

func modify(arr *[]int) {
	(*arr)[0] = 1
}

func modify2(arr []int) {
	arr[0] = 2
}
