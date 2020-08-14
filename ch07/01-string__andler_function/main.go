package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1.Contains 字符串str中是否包含substr，返回bool值
	var str string = "helllogooo"
	fmt.Println(strings.Contains(str, "go"))
	fmt.Println(strings.Contains(str, "asb"))

	// 2.join 字符串链接，把slicea通过sep链接起来
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "|")
	fmt.Println("buf= ", buf)

	// 3.Index 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("hellogo", "go"))
	fmt.Println(strings.Index("hellogo", "python"))

	// 4.Repeat 重复s字符串count次，最后返回重复的字符串
	bu := strings.Repeat("go", 3)
	fmt.Println("bu=", bu)

	// 5.Replace 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("abcgefgejke", "e", "python", 2))
	fmt.Println(strings.Replace("abcgefgejke", "e", "python", -1))

	// 6.Split 把s字符串按照sep分割，返回slice
	bu2 := "hello@abc@go@mike"
	s2 := strings.Split(bu2, "@")
	fmt.Println("s2=", s2)

	// 7.Trim 在s字符串的头部和尾部去除cutset指定的字符串
	bu3 := strings.Trim("    are you ok      ", " ")
	fmt.Printf("#%s#\n", bu3)

	// 8.Fields 去除s字符串的空格符，并且按照空格分割返回slice
	s3 := strings.Fields("    how are you    ")
	fmt.Println(s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}

}
