package main

import (
	"fmt"
	"math"
)

type myFunc func(num float64) string

func main() {
	//定义切片，对其中的元素进行求平方根和求平方的运算
	arr := []float64{1, 9, 16, 25, 30}
	result := FilterSlice(arr, func(num float64) string {
		num = math.Sqrt(num)
		return fmt.Sprintf("%.2f", num)
	})
	fmt.Println(result)
}

//遍历切片，对其中的每个元素进行运算处理
func FilterSlice(arr []float64, f myFunc) []string {
	var result []string
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}
