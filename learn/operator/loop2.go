package main

import "fmt"

func main() {
	//conditionFro()
	//traverseString()
	rraverseSlice()
}

func conditionFro() {
	result := 8
	for result > 5 {
		if result < 10 {
			fmt.Println(result)
			result++
		} else {
			break
		}
	}
}

//for ...range  对字符串、slice、数组、map等进行操作
func traverseString() {
	str := "ABCabc"
	for i, value := range str {
		fmt.Printf("第%d位字符为%c \n", i, value)
	}
}

func rraverseSlice() {
	arr := []int{100, 200, 300}
	for i, value := range arr {
		fmt.Printf("第%d位元素为%d \n", i, value)
	}
}
