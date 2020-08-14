package main

import "fmt"

func Test() int {
	var num1 int = 5
	var num2 int = 3
	var sum int
	sum = num1 + num2
	return sum
}

func Test1() (sum int) {
	var num1 int = 5
	var num2 int = 3
	sum = num1 + num2
	return sum
}

func Test2() (sum int) {
	var num1 int = 5
	var num2 int = 3
	sum = num1 + num2
	return
}

// 返回多个值
func Test3() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}

// 发送邮件信息
func SendMsg() {
	fmt.Println("邮件已发送------------")

}

// 完成用户信息校验
func CheckInfo(userName string, userPwd string, userEmail string) (b bool) {
	if userName != "" && userPwd != "" && userEmail != "" {
		b = true
	} else {
		b = false
	}
	return
}

// 注册用户信息
func Register() {
	var flag bool = false
	flag = CheckInfo("admin", "1234", "kadycui@163.com")
	if flag {
		SendMsg()
		fmt.Println("用户注册成功")
	} else {
		fmt.Println("用户注册失败")
	}
}

func main() {
	//var result int
	//result = Test2()
	//fmt.Println(result)

	//var r1 int
	//var r2 int
	//var r3 int
	//r1, r2, r3 = Test3()
	//fmt.Printf("r1=%d,r2=%d,r3=%d", r1, r2, r3)

	Register()
}
