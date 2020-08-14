package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomFunc() [3]int {
	var data [3]int                  // 定义数组,存储产生的随机数
	rand.Seed(time.Now().UnixNano()) // 以当前系统的时间作为种子参数
	for i := 0; i < 3; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	return data

}

func GetAvg() float64 {
	var sum int
	var numbers [3]int
	numbers = RandomFunc()
	for i := 0; i < len(numbers); i++ {
		fmt.Println("随机数:", numbers[i])
		sum += numbers[i]
	}
	return float64(sum) / float64(len(numbers))
}

func main() {
	//var num [3]int = [3]int{1, 2, 7}
	avg := GetAvg()
	fmt.Printf("%.2f", avg)
}
