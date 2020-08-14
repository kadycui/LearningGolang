package main

import "fmt"

func main() {
	// 定义1 初始化切片
	//s := []int{1, 2, 3}

	// 定义2 空切片
	//var s1 []int

	// 定义3 make()函数实现
	s := make([]int, 5, 10) // 切片的长度5 容量10
	fmt.Println(s)
	fmt.Println("切片长度:", len(s))
	fmt.Println("切片容量:", cap(s))

	// 切片赋值
	s[0] = 0
	s[1] = 1

	for i := 2; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println(s)

}
