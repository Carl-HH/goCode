package main

import "fmt"

/**
tips:
	1.一个函数最多只能有一个可变参数
	2.参数列表中还有其他类型的参数，则可变参数写在所有参数的最后
*/
func main() {
	//传进n个参数
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))
	//传切边作为参数
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8}...))
	sum, avg, count := getScore(10, 20, 30, 40, 50)
	fmt.Printf("sum: %.2f , avg: %.2f, count :%d \n", sum, avg, count)
	scores := []float64{10, 20, 30, 40, 50}
	sum, avg, count = getScore(scores...)
	fmt.Printf("sum: %.2f , avg: %.2f, count :%d", sum, avg, count)

}

//累加求和，参数个数不定，参数个数从0到n
func sum(numbs ...int) (sum int) {
	for _, value := range numbs {
		sum += value
	}
	return
}

//累加求总分及求平均分
func getScore(numbs ...float64) (sum, avg float64, count int) {
	for _, value := range numbs {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}
