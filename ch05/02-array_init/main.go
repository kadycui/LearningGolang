package main

import "fmt"

func main() {
	// 1.全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 2.部分初始化, 没有初始化的元素 赋值0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	// 3.指定元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)

	e := [...]int{1, 2, 3}
	fmt.Println(e[0])

}
