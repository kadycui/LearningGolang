package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("请输入出生日期,格式为:年-月-日(eg:2020-04-10)\n")
	fmt.Scan(&str)
	s := strings.Split(str, "-")
	fmt.Println(s[0] + "年" + s[1] + "月" + s[2] + "日")
}
