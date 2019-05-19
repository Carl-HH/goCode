package main

import "fmt"

func main() {
	breakcontinue()
}

func breakcontinue() {
	fmt.Println("\n1.break 语句")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}

	fmt.Println()

	fmt.Println("\n2.continue 语句")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
}
