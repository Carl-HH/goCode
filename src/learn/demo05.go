package main

/**
for：for 的三个组件都可省略，Go 没有 while，用 for 来替代
*/
import (
	"bufio"
	"fmt"
	"os"
)

func sum() int {
	var value int
	for i := 0; i <= 100; i++ {
		value += i
	}
	return value
}

// 等同于 while(true)
func deadLoop() {
	for {
		fmt.Println("this is a deadLoop")
	}
}

// Go 没有while，循环全部用 for，for的三个组件都可以省略
func printFile(filename string) {
	//xcode-select -- install
	// 打开文件
	file, err := os.Open(filename)
	// 如果出错，结束进程
	if err != nil {
		panic(err)
	}
	// 获取读取器
	scanner := bufio.NewScanner(file)
	// 读取：It returns false when the scan stops, either by reaching the end of the input or an error
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	printFile("abc.txt")
}

/**
for, if 后面都条件没有括号
if条件里也可以定义变量
没有while
switch不需要break，也可以直接switch多个条件
*/
