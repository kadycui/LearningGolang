package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) SayHello() {

	fmt.Println("我是人类,我叫:" + p.name)
}

type Chinese struct {
	Person
}

func (c *Chinese) SayHello() {

	fmt.Println("我是中国人,我叫:" + c.name)
}

type English struct {
	Person
}

func (e *English) SayHello() {

	fmt.Println("我是英国人,我叫:" + e.name)
}

type Personer interface {
	SayHello()
}

func main() {
	c := &Chinese{Person{"zhangsan"}}
	e := &English{Person{"MrZhang"}}
	x := make([]Personer, 2)
	x[0] = c
	x[1] = e
	for i := 0; i < len(x); i++ {
		x[i].SayHello()
	}
}
