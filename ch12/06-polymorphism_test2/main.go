package main

import "fmt"

type CarTyper interface {
	GetCar()
}
type CarStore struct {
}

func (c *CarStore) Order(money float64, carType string) {
	switch carType {
	case "BMW":
		CarWho(&BMWCar{Car{carType, money}})
	case "Audi":
		CarWho(&AudiCar{Car{carType, money}})
	}
}

type Car struct {
	carType string
	money   float64
}

func (p *Car) Move() {
	fmt.Println(p.carType + "汽车移动")
}
func (p *Car) Stop() {
	fmt.Println(p.carType + "汽车停止")
}

type BMWCar struct {
	Car
}

func (p *BMWCar) GetCar() {
	if p.money >= 60000 {
		p.Move()
		p.Stop()
	} else {
		fmt.Println("钱不够，无法买宝马！！")
	}
}

type AudiCar struct {
	Car
}

func (p *AudiCar) GetCar() {
	if p.money >= 70000 {
		p.Move()
		p.Stop()
	} else {
		fmt.Println("钱不够，无法买奥迪！！")
	}
}
func CarWho(i CarTyper) {
	i.GetCar()
}
func main() {
	var carStore CarStore
	carStore.Order(30000, "Audi")
}
