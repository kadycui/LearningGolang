package main

import "fmt"

func main() {
	names := [...]string{"我", "是", "好人"}

	var temp string

	for i := 0; i < len(names)/2; i++ {
		temp = names[i]
		names[i] = names[len(names)-1-i]
		names[len(names)-1-i] = temp
	}
	// 输出交换后数组中的值
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
