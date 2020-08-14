package main

import "fmt"

func PlayGame() {
	fmt.Print("1游戏开始!!!!!!!!!")
	fmt.Print("2游戏开始!!!!!!!!!")
	fmt.Print("3游戏开始!!!!!!!!!")
	fmt.Print("4游戏开始!!!!!!!!!")
	fmt.Print("5游戏开始!!!!!!!!!\n")
}

func Attack() {
	fmt.Println("远程攻击!!!")
	fmt.Println("拳头攻击!!!")
	fmt.Println("闪躲!!!\n")
}

// 普通参数
func SumAdd(a int, b int) {
	var sum int
	sum = a + b
	fmt.Println("两数的和为:", sum)
}

// 不定参数
func Test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
	fmt.Println("")
}

func Test1(a int, args ...int) {
	for i, data := range args {
		fmt.Println("编号为:", i)
		fmt.Println("值为:", data)
	}
}

// 在对函数进行调用时，固定参数必须传值，不定参数可以根据需要来决定是否要传值
func Test2(a int, b int, args ...int) {
	for i, data := range args {
		fmt.Println("编号为:", i)
		fmt.Println("值为:", data)
	}
	fmt.Printf("a=%d", a)
	fmt.Printf("b=%d", b)
}

func Test3(args ...int) {
	var sum int
	for _, data := range args {
		sum += data
	}
	fmt.Println("所有整数的和为:", sum)
}

func main() {
	//var num1, num2 int
	//fmt.Println("请输入第一个数:")
	//fmt.Scanf("%d\n",&num1)
	//fmt.Println("请输入第二个数:")
	//fmt.Scanf("%d\n",&num2)
	//SumAdd(num1,num2)

	//Test(1)
	//Test(1, 2)
	//Test(1, 2, 3)

	//Test1(1)
	//Test1(1, 2)
	//Test1(1, 2, 3)

	//Test2(1, 2, 3, 4, 5)

	Test3(3, 5, 9)

}
