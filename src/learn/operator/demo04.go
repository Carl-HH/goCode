package main

import "fmt"

/*
 数据类型转换
*/
func main() {
	chinese := 90
	english := 80.5
	avg := (float64(chinese) + english) / 2
	fmt.Println(avg)

	flag := true //布尔型无法和其他类型进行转换
	fmt.Println(flag)

	//str := "huachen"
	// int(str)
	//
	result := string(chinese)
	fmt.Println(result)

	chr := 'Z'
	result2 := int(chr)
	fmt.Println(result2)

}
