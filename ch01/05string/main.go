package main

import (
	"fmt"
)

/*
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠
*/

func main() {
	path := "\"D:\\Go\\src\\kadycui.com\\mygo\\day01\""
	fmt.Println(path)

	s := "I'm OK"
	fmt.Println(s)

	c := `
	白日依山尽
	huangheruhail
	欲穷千里目
	更上一层楼`
	fmt.Println(c)

	p := `D:\Go\src\kadycui.com\mygo\day01`
	fmt.Println(p)
}
