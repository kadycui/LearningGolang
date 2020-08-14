package main

import "fmt"

func main() {

	//var a [10]int

	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}
	//for i := 0; i < 10; i++{
	//	fmt.Println(a[i])
	//}

	//for i, data := range a {
	//	fmt.Printf("a[%d]=%d\t", i, data)
	//}

	for _, data := range a {
		fmt.Println("元素值:", data)
	}

}
