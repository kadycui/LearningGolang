package main

import "fmt"

func Swap(num1, num2 int) {
	num1, num2 = num2, num1
	fmt.Printf("num1=%d,num2=%d\n", num1, num2)
}

func PointerSwap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
	fmt.Printf("num1=%d,num2=%d\n", *num1, *num2)
}

func main() {

	var p *int
	p = new(int)
	// 通过new函数来指向
	// new(int)作用就是创建一个整型大小的空间
	//然后让指针变量p指向了该空间，所以通过指针变量p进行赋值后，该空间中的值就是57.
	*p = 57
	fmt.Println(*p)

	// 自动推导类型
	q1 := new(int)
	*q1 = 789
	fmt.Println(*q1)

	a := 10
	b := 20

	Swap(a, b)
	fmt.Printf("a=%d,b=%d", a, b)

	var c int = 30
	var d int = 40

	PointerSwap(&c, &d)
	fmt.Printf("c=%d,d=%d", c, d)

}
