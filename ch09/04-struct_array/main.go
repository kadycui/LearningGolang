package main

import "fmt"

type Student struct {
	id    int
	name  string
	score float64
}

// 结构体指针变量作为形参
func Test(p *Student) {
	p.id = 88
}

func main() {
	// 定义结构体指针 取结构体地址
	var p *Student = &Student{1, "Yaoming", 90}
	fmt.Println(*p)

	Test(p) // 传递结构体地址
	fmt.Println(*p)

	p1 := &Student{2, "Liuxiang", 93}
	p1.score = 88
	fmt.Println(*p1)

}
