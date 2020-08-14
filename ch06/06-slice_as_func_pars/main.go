package main

import "fmt"

func InitData(num [10]int) {
	for i := 0; i < len(num); i++ {
		num[i] = i
	}
}

// 切片是引用传递
// 数组是值传递

func main() {
	//s := make([]int, 10)   // 引用传递：也称为传地址。函数调用时，实际参数的引用(地址，而不是参数的值)被传递给函数中相对应的形式参数（实参与形参指向了同一块存储区域），在函数执行中，对形式参数的操作实际上就是对实际参数的操作，方法执行中形式参数值的改变将会影响实际参数的值。
	var s [10]int // 值传递：方法调用时，实参数把它的值传递给对应的形式参数，方法执行中形式参数值的改变不影响实际参数的值
	InitData(s)

	for _, v := range s {
		fmt.Println(v)
	}
}
