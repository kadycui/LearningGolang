package main

import "fmt"

// 数组指针
func Swap(p *[5]int) {
	(*p)[0] = 89
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	Swap(&a)
	fmt.Println(a)

	// 定义一个指针数组
	// 指针数组的定义方式，与数组指针定义方式是不一样的，注意指针数组是将“*”放在了下标的后面。
	var p [2]*int

	var i int = 20
	var j int = 30

	p[0] = &i
	p[1] = &j

	for index, value := range p {
		fmt.Printf("index=%d,value=%d\n", index, value)

	}
}
