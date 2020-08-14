package main

import "fmt"

// defer ⽤于延迟一个函数（或者当前所创建的匿名函数）的执行。注意，defer语句只能出现在函数的内部

func TestD(x int) {
	result := 100 / x
	fmt.Println(result)
}

// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行
func main() {
	//defer fmt.Println("这是延迟调用")
	//fmt.Println("正是正常调用")
	defer fmt.Println("你好")
	defer fmt.Println("世界")
	defer TestD(2)
	defer fmt.Println("Golang")

}
