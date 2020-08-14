package main

import "fmt"

// 定义接口
type Humaner interface {
	sayhi()
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

// 实现多态

// 定义一个普通函数，函数的参数为接口 只有一个函数，可以有不同的表现  多态
func Whosayhi(i Humaner) {
	i.sayhi()
}

func main() {
	var i Humaner
	s := &Student{"mike", 25}
	i = s
	i.sayhi()

	t := &Teacher{"Beijing", "go"}
	i = t
	i.sayhi()

}
