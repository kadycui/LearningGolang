package main

import "fmt"

// 存储的数据（内存），对应的地址

func main() {
	//var i int
	//i = 100
	//fmt.Printf("i=%d\n", i)
	//fmt.Printf("&i=%v\n", &i)  // i在内存中的地址

	var i int = 100
	var p *int // 定义一个指针变量
	p = &i

	*p = 80 // 使用指针修改变量
	fmt.Printf("i=%d,p=%v", i, p)

}
