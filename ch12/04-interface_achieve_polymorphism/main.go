package main

import "fmt"

type Flyabler interface {
	Fly()
}
type Bird struct {
}

func (b *Bird) EatAndDrink() { //为Bird定义方法
	fmt.Println("鸟儿吃喝")
}

type MaQue struct { //麻雀
	Bird
}

func (m *MaQue) Fly() {
	fmt.Println("麻雀会飞")
}

type YingWu struct {
	Bird
}

func (y *YingWu) Fly() {
	fmt.Println("鹦鹉飞")
}

type Plane struct {
}

func (p *Plane) Fly() {
	fmt.Println("飞机飞")
}
func WhoFly(i Flyabler) {
	i.Fly()
}

func main() {
	m := &MaQue{}
	m.EatAndDrink()
	WhoFly(m)

	plane := &Plane{}
	WhoFly(plane)
}
