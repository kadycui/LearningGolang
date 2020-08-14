package main

import "fmt"

func Test(a int, b int) (sum int) {
	sum = a + b
	return sum
}

type FuncType func(a int, b int) int // 定义函数类型

var a int

/*
 函数作用域
(1)在函数外边定义的变量叫做全局变量。
(2)全局变量能够在所有的函数中进行访问
(3)如果全局变量的名字和局部变量的名字相同，那么使用的是局部变量的
*/
func Test1() {
	a = 5
	a += 1
	fmt.Println("Test", a)
}

func main() {
	//var s int
	//var r FuncType
	//r = Test
	//s = r(3, 8)
	//fmt.Println(s)

	a := 9
	Test1()
	fmt.Println("main", a)

}
