package main

import "fmt"

type Operation struct { //定义操作父类
	numA float64
	numB float64
}
type GetResulter interface { //定义接口
	GetResult() float64
}
type OperationAdd struct { //加法
	Operation
}

func (a *OperationAdd) GetResult() float64 { //实现接口
	return a.numA + a.numB
}

type OperationSub struct { //减法
	Operation
}

func (s *OperationSub) GetResult() float64 {
	return s.numA - s.numB
}

//创建一个类负责对象创建
type OperationFactory struct {
}

func (f *OperationFactory) CrateOption(option string, numA float64, numB float64) float64 {
	var result float64
	switch option {
	case "+":
		add := &OperationAdd{Operation{numA, numB}}
		result = OperationWho(add)
	case "-":
		sub := &OperationSub{Operation{numA, numB}}
		result = OperationWho(sub)
	}
	return result
}
func OperationWho(i GetResulter) float64 {
	return i.GetResult()
}

func main() {
	var opfactory OperationFactory
	s := opfactory.CrateOption("-", 10, 12)
	fmt.Println(s)
}
