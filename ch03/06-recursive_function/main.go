package main

import "fmt"

func Test(a int) {
	if a == 1 {
		fmt.Println("a=", a)
		return
	}
	Test(a - 1)
	fmt.Println("abc=", a)
}

func main() {
	Test(3)
}
