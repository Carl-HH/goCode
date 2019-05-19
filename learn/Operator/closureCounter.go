package main

import "fmt"

//闭包的另外一种写法
func main() {
	res := func() func() int {
		i := 10
		return func() int {
			i++
			return i
		}
	}()
	fmt.Println(res())
}
