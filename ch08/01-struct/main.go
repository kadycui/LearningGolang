package main

import "fmt"

func main() {
	// 定义结构体
	type Student struct {
		id   int
		name string
		sex  byte
		age  int
		addr string
	}

	// 初始化
	var s1 Student = Student{1, "mike", 'm', 15, "Beijing"}
	fmt.Println("s1=", s1)

	s2 := Student{name: "apple", addr: "Shanghai"}
	fmt.Println("s2=", s2)

	// 使用
	var s Student // 定义结构体变量
	s.id = 1001
	s.name = "旺旺"
	s.sex = 'w'
	s.age = 18
	s.addr = "Zhengzhou"
	fmt.Println(s)

	// 结构体的比较和和赋值
	s3 := Student{1, "mike", 'm', 18, "sh"}
	s4 := Student{1, "mike", 'm', 18, "sh"}
	s5 := Student{2, "mike", 'm', 18, "sh"}
	fmt.Println("s3==s4", s3 == s4)
	fmt.Println("s3==s5", s3 == s5)

	s6 := Student{3, "mike", 'm', 19, "sz"}
	var tmp Student
	tmp = s6
	fmt.Println(tmp)

}
