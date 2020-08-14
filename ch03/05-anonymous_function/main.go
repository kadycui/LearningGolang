package main

import "fmt"

func main() {
	//var num int
	//num = 9
	//f := func() {
	//	num++
	//	fmt.Println("匿名函数:", num)
	//}
	//f()
	//fmt.Println("main函数:", num)
	f := func(a, b int) {
		var sum int
		sum = a + b
		fmt.Println("和为:", sum)
	}
	f(85, 96)

	// 匿名函数,有参有返回值
	//x, y := func(i, j int) (max, min int) {
	//	if i > j {
	//		max = i
	//		min = j
	//	} else {
	//		max = j
	//		min = i
	//	}
	//	return
	//}(10, 20)
	//fmt.Printf("x=%d,y=%d\n", x, y)
}
