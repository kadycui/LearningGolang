package main

import "fmt"

// 匿名字段

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person

	score float64
}

func main() {
	var s1 Student = Student{Person{101, "Tom", 19}, 99.5}
	fmt.Println("s1=", s1)

	// 自动推导初始化
	s2 := Student{Person{102, "Jim", 20}, 98.1}
	fmt.Printf("s2= %+v\n", s2) // %+v 详细显示
	// 直接给s2中的Person成员(匿名字段)赋值
	s2.Person = Person{112, "Ayinst", 80}
	fmt.Printf("s2= %+v\n", s2)

	// 指定初始化成员
	s3 := Student{score: 99.1}
	fmt.Printf("s3= %+v\n", s3)

	// 成员操作
	var s4 Student = Student{Person{104, "Nigols", 25}, 88.5}
	s4.age = 30
	s4.Person.id = 120
	s4.score = 100
	fmt.Println(s4)
}
