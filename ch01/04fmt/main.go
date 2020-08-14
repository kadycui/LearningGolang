package main

import "fmt"

func main() {
	var n = 100

	// 查看类型
	fmt.Printf("%T\n", n) // 类型
	fmt.Printf("%v\n", n) // 二进制
	fmt.Printf("%b\n", n) //
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello 你好!"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}
