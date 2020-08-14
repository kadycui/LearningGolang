package main

import "fmt"

func InitData(num []int) {
	for i := 0; i < len(num); i++ {
		fmt.Printf("请输入第%d个数:\n", i+1)
		fmt.Scanf("%d\n", &num[i])
	}
}

func SumAdd(num []int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func main() {
	var count int
	var sum int
	fmt.Println("请输入整数个数:")
	fmt.Scanf("%d", &count)
	s := make([]int, count)
	InitData(s)
	sum = SumAdd(s)
	fmt.Println("Sum=", sum)
}
