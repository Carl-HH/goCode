package main

import "fmt"

func main() {
	loop1()
}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i)
	}

	fmt.Println()

	//i的作用域不同
	i := 0
	for ; i < 10; i++ {
		fmt.Print(" ", i)
	}

	fmt.Println()

	for ; ; i++ {
		if i < 20 {
			fmt.Print(" ", i)
		} else {
			break
		}
	}

	fmt.Println()

	for { //同 for {}
		if i < 30 {
			fmt.Print(" ", i)
			i++
		} else {
			break
		}
	}

}
