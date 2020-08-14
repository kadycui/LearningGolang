package main

import "fmt"

type Humaner interface { // 子集
	sayhi()
}

type Personer interface { // 超集
	Humaner // 匿名字段， 继承sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

// Student实现了sayhi()
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱： ", lrc)
}

func main() {
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 23}
	i = s
	i.sayhi() // 继承过来的方法
	i.sing("好汉歌")

	// 超集可以转换成子集
	var iPro Personer // 超集
	iPro = &Student{"mikepro", 25}

	var j Humaner
	j = iPro
	j.sayhi()

}
