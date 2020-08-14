package main

import "fmt"

// 定义结构体
type Student struct {
	name    string
	sex     byte
	age     int
	chinese float64
	math    float64
	english float64
}

// 定义结构体方法
func (p *Student) SayHello(name string, age int, sex byte) {
	p.name = name
	p.age = age
	p.sex = sex

	// 年龄性别进行判断
	if p.age < 0 && p.age > 100 {
		p.age = 0
	}
	if p.sex != 'm' && p.sex != 'w' {
		p.sex = 'm'
	}
	fmt.Printf("我叫%s，我今年%d岁了，我是%c生\n", p.name, p.age, p.sex)
}

func (p *Student) ShowScore(chinese float64, math float64, english float64) {
	p.chinese = chinese
	p.math = math
	p.english = english

	var sum float64
	sum = p.chinese + p.math + p.english
	fmt.Printf("我叫%s,我的总成绩是%.1f,平均分是%.1f", p.name, sum, sum/3)
}

func main() {
	var stu Student
	stu.SayHello("Xiaoxiao", 19, 'm')
	stu.ShowScore(90, 85, 79)
}
