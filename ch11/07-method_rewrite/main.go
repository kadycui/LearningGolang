package main

import "fmt"

//type Person struct {
//	name   string
//	age    int
//	gender byte
//}
//
//type Student struct {
//	Person
//	id   int
//	addr string
//}
//
//func (p *Person) PrintInfo() {
//	fmt.Printf("姓名=%s,性别=%c,年龄=%d\n", p.name, p.gender, p.age)
//}
//
//// 方法重写
//func (tmp *Student) PrintInfo() {
//	fmt.Println("Student: tmp=", tmp)
//}

//func main() {
//	s := Student{Person{"Zhangsan", 19, 'm'}, 1, "北京"}
//	s.PrintInfo() // 调用Student
//
//	// 显示调用
//	s.Person.PrintInfo()
//}

type Animal struct {
	age int
}

func (p *Animal) Bark() {
	fmt.Println("汪汪叫")
}

type Dog struct {
	Animal
	name string
}
type Cat struct {
	Animal
	name string
}

func (c *Cat) Bark() {
	fmt.Println("喵喵叫")
}

func main() {
	var dog Dog
	dog.Bark()

	var cat Cat
	cat.Bark()
}
