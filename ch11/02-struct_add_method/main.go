package main

import "fmt"

// 定义一个结构体
type Student struct {
	id    int
	name  string
	age   int
	score float64
}

//    方法的接收者
func (stu Student) PrintShow(id int, name string, age int, score float64) {
	fmt.Println(stu)
}

// 接收者为指针类型
func (p *Student) EditInfo(id int, name string, age int, score float64) {
	p.id = id
	p.name = name
	p.age = age
	p.score = score
	fmt.Println("EditInfo:", *p)
}

// 接收者为普通变量，非指针，值传递
// 接收者为指针变量，引用传递

func main() {
	s := Student{100, "Lisy", 21, 88}
	s.PrintShow(100, "Lisy", 21, 88) // 创建完对象调用方法

	var stu Student
	// 将结构体的指针传递给接收者，同时将要修改的数据传递到方法中
	(&stu).EditInfo(1002, "Zhangsan", 19, 89)
	stu.PrintShow(100, "Lisy", 21, 88)

	var stu2 Student
	(&stu2).PrintShow(101, "Qianqi", 24, 96)
	(&stu2).EditInfo(1033, "Sunba", 23, 65)

	var stu3 Student
	stu3.EditInfo(109, "Zhoujiu", 25, 94) // 将普通的结构体类型的变量转换成(&stu)在调用EditInfo( )方法
}
