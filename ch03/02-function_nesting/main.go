package main

import "fmt"

func Test1(num1 int, num2 int) {
	fmt.Println(num1 + num2)
	fmt.Println("Test1")
}

// 基本函数嵌套
func Test(a int, b int) {
	Test1(a, b)
	fmt.Println("Test")
}

// 不定参数函数嵌套
func Test2(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}

}

func Test3(args ...int) {
	Test2(args[2:]...)
}

func main() {
	//Test(3,9)
	//fmt.Println("main")
	Test3(3, 5, 9)
}
