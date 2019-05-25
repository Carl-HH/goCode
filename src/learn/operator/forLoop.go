package main

import "fmt"

//for嵌套循环语句
func main() {
	forLoop7()
}

func forLoop() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i: %d , j : %d \n", i, j)
		}
	}
}

func forLoop3() {
	for i := 0; i < 5; i++ {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func forLoop4() {
	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func forLoop5() {
	for i := 5; i > 0; i-- {
		for j := 0; j <= 5; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func forLoop6() {
	for i := 1; i < 5; i++ {
		for m := 5; m >= i; m-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func forLoop7() {
	for i := 0; i < 5; i++ {
		for j := i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for m := 5; m > i; m-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
