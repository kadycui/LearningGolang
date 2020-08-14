package main

import "fmt"

type Object struct {
	id   int
	flag bool
}

type Person struct {
	Object
	name string
	age  int
}

type Student struct {
	Person
	name  string
	score float64
}

func main() {
	var s Student
	s.name = "Lisi"
	s.Person.name = "Wangwu"
	s.Person.Object.id = 110
	fmt.Printf("s1 = %+v\n", s)
}
