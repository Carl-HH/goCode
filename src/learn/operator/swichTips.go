package main

import "fmt"

func main() {
	eval2()
}

func eval() {
	num1, num2, result := 12, 4, 0
	operation := "+"
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		result = -1
	}
	fmt.Println(result)
}

//fallthrough
//强制执行后面的一个case分支，必须放在case分支的最后一行。
func eval2() {
	num1, num2, result := 12, 4, 0
	operation := "+"
	switch operation {
	case "+":
		result = num1 + num2
		fallthrough
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2

	default:
		result = -1

	}
	fmt.Println(result)
}
