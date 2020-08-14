package main

import "fmt"

// 方法值与方法表达式
type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}
func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 25}
	fmt.Printf("main: %p,%v\n", &p, p)

	// 方法值  f := p.SetInfoPointer   //隐藏了接收者

	// 方法表达式
	f := (*Person).SetInfoPointer
	f(&p) // 显式把接受者传递过去 ---->   p.SetInfoPointer()

	f2 := (Person).SetInfoValue
	f2(p) // // 显式把接受者传递过去 ---->   p.SetInfoValue()
}
