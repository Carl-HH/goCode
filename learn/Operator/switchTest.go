package main

import "fmt"

func main() {
	//defaultSwitch()
	//score()
	mouthDay()
}

//switch 默认每个分支自带break
func defaultSwitch() {
	switch { //省略表达式，默认为true
	case true:
		println("true")
	case false:
		println("false")
	}
}

func score() {
	score := 75
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	case score >= 0:
		grade = "E"
	}
	println("score:", score)
	println("grade:", grade)
}

func mouthDay() {
	year := 2000
	month := 2
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		day = -1
	}
	fmt.Printf("%d-%d-%d", year, month, day)
}
