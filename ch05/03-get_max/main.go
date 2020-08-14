package main

import "fmt"

func main() {
	// 从一个整数数组中取出最大的整数,最小整数,总和,平均值
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	// 声明两个变量存储最大值和最小值
	var max int = a[0]
	var min int = a[0]
	var sum int

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		sum += a[i]
	}
	fmt.Printf("最大值是:%d,最小值是:%d,和为:%d,平均值为:%d", max, min, sum, sum/len(a))
}
