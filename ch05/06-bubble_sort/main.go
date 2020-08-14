package main

import "fmt"

func main() {
	// 冒泡排序
	var num [10]int = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var temp int
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				temp = num[j]
				num[j] = num[j+1]
				num[j+1] = temp
			}
		}
	}
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}
