package main

import "fmt"

//打印格式化

type point struct {
	x, y int
}

func main() {
	p := point{10, 20}
	fmt.Printf("%T , %v \n", p, p)

	fmt.Printf("%T , %t \n", true, true)

	fmt.Printf("%T , %d \n", 5, 5)
	fmt.Printf("%T , %5d \n", 123, 123)

	fmt.Printf("%T , %b \n", 123, 123)

	value := fmt.Sprintf("%b", 123)
	fmt.Println(value)

	fmt.Printf("%T , %x \n", 123, 123)
	fmt.Printf("%T , %X \n", 123, 123)
	fmt.Printf("%T , %U \n", '一', '一')

	//浮点型
	fmt.Printf("%f \n", 6.5) //默认小数点后5位
	fmt.Printf("%.2f \n", 6.5)
	fmt.Printf("%.10e \n", 126.534535)

	arr := []byte{97, 98, 99, 65}
	arr2 := []rune{'a', 'b', 'c', 'A'}
	//字符串
	fmt.Printf("%s \n", "你好")
	fmt.Printf("%q \n", "你好")

	fmt.Printf("%T , %s \n", arr, arr)
	fmt.Printf("%T , %s \n", arr2, arr2)
	fmt.Printf("%T , %x \n", arr, arr)
	fmt.Printf("%T , %X \n", arr, arr)

}
