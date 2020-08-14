package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

// 结构体作为函数参数
func UseStruct(stu Student) {
	stu.id = 6
	fmt.Println("stu:", stu)
}

func main() {
	s := Student{1, "mike", 'm', 19, "zz"}
	UseStruct(s)
	fmt.Println("s=", s)
}
