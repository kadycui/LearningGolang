package main

import "fmt"

func main() {
	country := [...]string{"美利坚", "法兰西", "英吉利"}

	var str string

	for i := 0; i < len(country)-1; i++ {
		str += country[i] + "|"
	}
	fmt.Println(str + country[len(country)-1]) // 美利坚|法兰西|英吉利
}
