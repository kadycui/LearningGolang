package main

import (
	"fmt"
)

// 声明变量,推荐使用小驼峰命名

// var student_name string    （不推荐）
// var studentName string
// var StudentName string      （不推荐）

// var name string
// var age int
// var isOk bool

// 批量声明变量
var (
	name string
	age  int
	isOK bool
)

// 常量声明
// const pi = 301415926
// const e = 2.7182

// 批量声明常量
const (
	pi = 3.1415926
	e  = 207182
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

/*
iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
(iota可理解为const语句块中的行索引)。
 使用iota能简化定义，在定义枚举时很有用。
*/
const (
	a1 = iota //0
	a2        //1
	a3        //2
	a4        //3
)
const (
	b1 = iota //0
	b2 = iota //1
	_  = iota
	b3 = iota //3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1 1,d2 2
	d3, d4 = iota + 1, iota + 2 // d3 2 ,d4 3
)

const (
	a, b  = iota + 1, iota + 2 //1,2
	c, d                       //2,3
	ee, f                      //3,4
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func foo() (int, string) {
	return 10, "111111111"
}

func main() {
	name = "张三"
	age = 25
	isOK = true
	// 声明的变量必须使用,且不能重复声明

	fmt.Printf("name:%s", name)
	fmt.Println(age)
	fmt.Print(isOK)

	var address string = "广东广州"
	fmt.Println(address)

	// 简短变量声明,只能在函数里面使用
	s3 := "哈哈哈哈哈"
	fmt.Println(s3)

	// 匿名变量使用一个 _ 表示
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Println("常量π的值", pi)
	fmt.Println("常量e的值", e)

	fmt.Println("声明多个变量值n1,n2,n3 =", n1, n2, n3)

	fmt.Println("在const关键字出现时将被重置为0:", a1, a2, a3, a4)

	fmt.Println("在const关键字出现时将被重置为0:", b1, b2, b3)
	// println("d1", d1)
	// println("d2", d2)
	// println("d3", d3)
	// println("d4", d4)

	// println("a", a)
	// println("b", b)
	// println("c", c)
	// println("d", d)
	// println("ee", ee)
	// println("f", f)

	println("KB", KB)
	println("MB", MB)
	println("GB", GB)
	println("TB", TB)
	println("PB", PB)

}
