package main

import "fmt"

func main() {
	var b byte = 100 //(代表ASCII中的一个字符)
	var c rune = 200 //代表UTF-8 中的一个字符
	var e byte = 'a'
	var f rune = '一'
	var str string = "abc"

	fmt.Printf("%T , %v \n", b, b)
	fmt.Printf("%T , %v \n", c, c)
	fmt.Printf("%T , %v \n", e, e)
	fmt.Printf("%T , %v \n", f, f)
	fmt.Printf("%T , %v \n", str, str)

	//"`"用于多行字符串的输出
	temp := `	m := "19"
	n := "190"
	m, n, q := "a", "b", 10`
	fmt.Println(temp)
}
