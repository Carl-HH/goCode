package main

import "fmt"

//常量
func main() {
	const NAME string = "Carl"
	const a = iota
	fmt.Println(a)

	const (
		A = 10
		B
		C
	)
	fmt.Println(A, B, C)

	const (
		X = iota
		Y
		Z
	)
	fmt.Println(X, Y, Z)

	//iota的特点，常量计算器，和iota所在位置无关
	const (
		L = "CARL"
		M = iota
		N
	)
	fmt.Println(L, M, N) //CARL 1 2

	const (
		i = 1 << iota
		j = 3 << iota
		k // k = 3 << iota
		l // l = 3 << iota
	)
	fmt.Println(i, j, k, l) //1 6 12 24

	const (
		a1 = '-'
		b1
		c1 = iota
		d1
	)
	fmt.Println(a1, b1, c1, d1) //19968 19968 2 3

}
