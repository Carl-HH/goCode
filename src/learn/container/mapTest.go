package main

import "fmt"

func main() {
	//map使用哈希表，必须可以比较相等
	//除slice，map，function的内建类型都可以作为key
	//struct结构不包含以上字段，也可以作为key
	m := map[string]string{
		"name":   "huachen",
		"course": "golang",
	}

	m2 := make(map[string]string) //m2 == empty map

	var m3 map[string]string // m3 == nil

	fmt.Println(m, m2, m3)

	//map遍历，不保证顺序
	for k, v := range m {
		fmt.Printf("k: %s , v: %s \n", k, v)

	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	corseNmae := m["corce"] //key不存在时，获取value类型的初始值
	fmt.Println(corseNmae)

	fmt.Println("Getting values")
	courseName2, ok := m["course"] //用value,OK:=m[key]来判断是否存在key
	fmt.Println(courseName2, ok)
	corseNmae2, ok := m["corce"]
	fmt.Println(corseNmae2, ok)

	if courseName3, ok := m["corse"]; ok {
		fmt.Println(courseName3)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m, "name") //删除key

}
