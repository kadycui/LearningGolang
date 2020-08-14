package main

import "fmt"

// 定义结构体
type Student struct {
	id    int
	name  string
	score float64
}

func main() {
	students := []Student{
		Student{
			101,
			"张三",
			100,
		},
		Student{
			102,
			"李四",
			96,
		},
		Student{
			103,
			"王五",
			91,
		},
	}
	fmt.Println(students)
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i].name)
	}
}
