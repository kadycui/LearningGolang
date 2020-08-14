package main

import "fmt"

/*
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false

*/

//定义计算器的接口
type CalcSuper interface {
	SetData(data ...interface{}) //获取数据，验证数据
	CalcOperate() float64        //完成计算
}
type Operation struct { //定义操作父类
	numA float64
	numB float64
}
type Add struct { //定义加法类，继承父类
	Operation
}

func NewAdd() *Add { //创建Add对象，返回指针类型
	instance := new(Add)
	return instance
}
func (a *Add) SetData(data ...interface{}) { //让Add类实现该方法，完成数据校验，数据获取
	if len(data) != 2 {
		fmt.Println("error,Needtwoparameters")
		return
	}
	if _, ok := data[0].(float64); !ok {
		fmt.Println("error,Needfloat64parameters")
		return
	}
	if _, ok := data[1].(float64); !ok {
		fmt.Println("error,Needfloat64parameters")
		return
	}
	a.numA, _ = data[0].(float64)
	a.numB, _ = data[1].(float64)
}
func (a *Add) CalcOperate() float64 { //实现该方法，完成数据加法
	return a.numA + a.numB
}

type Subtraction struct { //定义减法类
	Operation
}

func NewSubtraction() *Subtraction {
	instance := new(Subtraction)
	return instance
}
func (a *Subtraction) SetData(data ...interface{}) {
	if len(data) != 2 {
		fmt.Println("error,Needtwoparameters")
		return
	}
	if _, ok := data[0].(float64); !ok {
		fmt.Println("error,Needfloat64parameters")
		return
	}
	if _, ok := data[1].(float64); !ok {
		fmt.Println("error,Needfloat64parameters")
		return
	}
	a.numA, _ = data[0].(float64)
	a.numB, _ = data[1].(float64)
}
func (a *Subtraction) CalcOperate() float64 {
	return a.numA - a.numB
}

type CalcFactory struct { //解决对象创建问题
}

func NewCalcFactory() *CalcFactory {
	instance := new(CalcFactory)
	return instance
}
func (f *CalcFactory) CreateOperate(opType string) CalcSuper {
	var op CalcSuper
	switch opType {
	case "+":
		op = NewAdd()
	case "-":
		op = NewSubtraction()
	default:
		panic("error!donthasthisoperate")
	}
	return op
}
func main() {
	//调用
	factory := NewCalcFactory()
	op := factory.CreateOperate("+")
	op.SetData(1.5, 2.0)
	fmt.Println(op.CalcOperate())
	op = factory.CreateOperate("-")
	op.SetData(1.5, 2.0)
	fmt.Println(op.CalcOperate())
}
