package main

import "fmt"

type Ticket struct {
	distance float64
	price    float64
}

func (p *Ticket) GetValue(price float64, distance float64) {
	if distance < 0 {
		distance = 0
	}
	p.distance = distance
	p.price = price

	if p.distance > 0 && p.distance <= 100 {
		p.price = p.price * 1.0
	} else if p.price >= 101 && p.price < 200 {
		p.price = p.price * 0.95
	} else if p.price >= 201 && p.price < 300 {
		p.price = p.price * 0.9
	} else {
		p.price = p.price * 0.8
	}
	fmt.Println(p.price)
}

func main() {
	var ticket Ticket
	ticket.GetValue(190, 200)
}
