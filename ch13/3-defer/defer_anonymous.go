package main

import "fmt"

// defer与匿名函数结合使用

func main() {
	a := 10
	b := 20

	//defer func(){
	//	fmt.Println("匿名函数中a", a)
	//	fmt.Println("匿名函数中b", b)
	//}()

	defer func(a int, b int) {
		fmt.Println("匿名函数中a", a)
		fmt.Println("匿名函数中b", b)
	}(a, b) // 虽然没有立即调用执行匿名函数，但是已经完成了参数的传递

	a = 100
	b = 200
	fmt.Println("main函数中a", a)
	fmt.Println("main函数中b", b)
}

// 未调用参数结果 ()
/*
main函数中a 100
main函数中b 200
匿名函数中a 100
匿名函数中b 200
*/

// 调用参数 (a, b)
/*
main函数中a 100
main函数中b 200
匿名函数中a 10
匿名函数中b 20

*/
