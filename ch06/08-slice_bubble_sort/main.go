package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func InitData(s []int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) // 100以内的随机数
	}
}

func main() {
	n := 10
	s := make([]int, n) // 创建一个切片
	InitData(s)         // 初始化数组
	fmt.Println("排序前: ", s)
	BubbleSort(s)
	fmt.Println("排序后: ", s)
}
