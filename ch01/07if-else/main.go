package main

import "fmt"

func main() {
	fmt.Println("请输入考试成绩:")
	var score int
	fmt.Scanf("%d", &score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
