package main

import "fmt"

func main() {
	//var userName string
	//var userPwd string
	//var count int
	//
	//for {
	//	fmt.Println("姓名:")
	//	fmt.Scanf("%s\n", &userName)
	//	fmt.Println("密码:")
	//	fmt.Scanf("%s\n", &userPwd)
	//
	//	if userName == "admin" && userPwd == "123456" {
	//		fmt.Println("Success!")
	//		// break的作用就是跳出本次循环
	//		break
	//	}else {
	//		count++
	//		if count >= 3{
	//			fmt.Println("密码错误次数太多!")
	//			break
	//		}
	//		fmt.Println("请重新输入!!!!!")
	//	}
	//}

	var age int
	var sum int
	var b bool = true

	for i := 1; i <= 5; i++ {
		fmt.Printf("请输入第%d个人的年龄:\n", i)
		fmt.Scanf("%d\n", &age)

		if age < 0 || age > 100 {
			fmt.Println("输入有误!!!!")
			b = false
			break
		}
		sum += age
	}
	if b {
		fmt.Println("5个人年龄的总和是", sum)
	}
}
