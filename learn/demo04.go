package main

/**
if：变量可以定义在 if 块内，其作用域就只在 if 块内了
*/
import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "abc.txt"
	// Go 函数可以返回两个值
	// func ReadFile(filename string) ([]byte, error)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// contents 是 []byte, 用%s 可以打印出来
		fmt.Printf("%s", contents)
	}
	// if 语句外部可访问
	fmt.Printf("%s", contents)
}

func readFileShorter() {
	const filename = "abc.txt"
	// Go 函数可以返回两个值
	// func ReadFile(filename string) ([]byte, error)

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		// contents 是 []byte, 用%s 可以打印出来
		fmt.Printf("%s", contents)
	}
	// if 语句外部不可访问
	//fmt.Printf("%s", contents) // 报错
}

/**
switch：默认自带break，如果想穿下去执行，使用 fallthrough
*/
// switch 默认自带break，如果想穿下去执行，使用 fallthrough
func eval(a, b int, op string) int {
	var value int
	switch op {
	case "+":
		value = a + b
	case "-":
		value = a - b
	default:
		panic("unsupport operator" + op)
	}
	return value
}
func main() {
	fmt.Print(eval(1, 2, "+"))
}
