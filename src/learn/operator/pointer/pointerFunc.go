package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 5
	fmt.Printf("a = %d , b = %d", a, b)
	fmt.Println()
	change(&c)
	fmt.Println("c = ", c)
	swapPointer2(&a, &b)
	fmt.Printf("a = %d , b = %d", a, b)
	fmt.Println()
	b, a = swap(a, b)
	fmt.Printf("a = %d , b = %d", a, b)
}

func change(num *int) {
	*num = 20
}

func swapPointer(a, b *int) {
	a, b = b, a
}

func swapPointer2(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b int) (int, int) {
	return b, a
}
