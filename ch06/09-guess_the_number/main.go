package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机三位数数字
func InitData() int {
	var num int
	rand.Seed(time.Now().UnixNano())
	for {
		num = rand.Intn(1000)
		if num >= 100 && num <= 999 {
			break
		}
	}
	return num
}

// 获取三位数字中的每一位数字
func GetNum(randSlice []int, num int) {
	randSlice[0] = num / 100
	randSlice[1] = num % 100 / 10
	randSlice[2] = num % 10
}

func onGame(randSlice []int) {
	var num int
	keySclice := make([]int, 3)
	for {
		for {
			fmt.Println("请输入一个三位数:")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入的数据不符合要求!")
		}
		GetNum(keySclice, num)
		n := 0
		for i := 0; i < 3; i++ {
			if keySclice[i] > randSlice[i] {
				fmt.Printf("第%d位大了点\n", i+1)
			} else if keySclice[i] < randSlice[i] {
				fmt.Printf("第%d位小了点\n", i+1)
			} else {
				fmt.Printf("第%d猜对了\n", i+1)
				n++
			}
		}
		if n == 3 {
			fmt.Println("全部猜对!!")
			break
		}
	}
}

func main() {
	num := InitData()
	s := make([]int, 3)
	GetNum(s, num)
	onGame(s)
}
