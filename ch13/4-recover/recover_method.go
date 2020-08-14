package main

import "fmt"

func TestA() {
	fmt.Println("func TestA()")
}

// Go语言为我们提供了专用于“拦截”运行时panic的内建函数——recover
// 注意：recover只有在defer调用的函数中有效。

func TestB(x int) {

	// 设值 recover
	defer func() {
		//recover()
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() // 调用此函数
	var a [10]int
	a[x] = 222
}

func TestC() {
	fmt.Println("func TestA()")
}

func main() {
	TestA()
	TestB(14)
	TestC()
}
