package main

import "fmt"

// 同名字段
type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person
	name  string
	score float64
}

// 指针类型匿名字段
type Person2 struct {
	id   int
	name string
	age  int
}

type Student2 struct {
	*Person2 // 指针类型匿名字段
	score    float64
}

func main() {
	var s1 Student
	s1.name = "Zhangsan"
	fmt.Printf("s1=%+v\n", s1)

	var s2 Student2

	s2 = Student2{&Person2{101, "Lisa", 19}, 96}
	fmt.Println(s2)
	fmt.Println(s2.id, s2.name, s2.age, s2.score)

	var s3 Student2
	s3.Person2 = new(Person2)
	s3.id = 103
	s3.name = "Kiti"
	s3.score = 97
	fmt.Printf("s3 = %+v\n", s3)
}
