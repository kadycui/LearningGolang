package main

import (
	"fmt"
	"strconv"
)

//  Format&Parse&Append
func main() {
	// bool转换成字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)

	// 整型转字符串
	var str2 string
	str2 = strconv.Itoa(123)
	fmt.Println(str2)

	// 浮点转成字符串
	var str3 string
	str3 = strconv.FormatFloat(3.14, 'f', 3, 64) // f 指以小数方式, 3小数点位数, 64指64位处理器
	fmt.Println("str = ", str3)

	// 字符串转换成其他类型
	var flag bool
	var err error

	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag=", flag)
	} else {
		fmt.Println("err=", err)
	}

	// 把字符串转换成整型
	a, _ := strconv.Atoi("568")
	fmt.Println("a=", a)
	b, err2 := strconv.ParseFloat("125.23", 64)
	if err2 == nil {
		fmt.Println("flag=", b)
	} else {
		fmt.Println("err2=", err)
	}

	//转换字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)

	// 第二个为要追加的数, 第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 123456, 10)
	slice = strconv.AppendQuote(slice, "dashijdhasjkhk")
	fmt.Println("slice= ", string(slice))

}
