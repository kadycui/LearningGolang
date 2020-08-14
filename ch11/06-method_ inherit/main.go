package main

import (
	"fmt"
)

// 方法体的继承

// 定义公共类结构体
type Person struct {
	name   string
	age    int
	gender byte
}

// 为Person结构体定义方法,完成公共属性赋值
func (p *Person) SetValue(name string, age int, gender byte) {
	p.name = name
	p.age = age
	p.gender = gender
}

type Reporter struct {
	Person
	hobby string
}

func (r *Reporter) ReporterSayHello(h string) {
	r.hobby = h
	fmt.Printf("我叫%s,是一名记者,我的爱好是%s,我是%c生,我今年%d岁了\n", r.name, r.hobby, r.gender, r.age)
}

type Programmer struct {
	Person
	WorkeYear int
}

func (s *Programmer) ProgrammerSayHello(work int) {
	s.WorkeYear = work
	fmt.Printf("我叫%s,是一名程序员,我是%c生,我今年%d岁了,我的工作年限是%d年", s.name, s.gender, s.age, s.WorkeYear)
}

func main() {
	var reporter Reporter
	reporter.SetValue("张三", 26, 'm')
	reporter.ReporterSayHello("采访")

	var programmer Programmer
	programmer.SetValue("李四", 26, 'm')
	programmer.ProgrammerSayHello(3)
}
