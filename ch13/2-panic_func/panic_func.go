package main

import "fmt"

func TestA() {
	fmt.Println("func TestA()")
}

func TestB(x int) {
	var a [10]int
	a[x] = 222
}

func TestC() {
	fmt.Println("func TestA()")
}

func main() {
	TestA()
	TestB(11)
	TestC()
}
