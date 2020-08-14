package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "ch14/1-make_text/text/a.txt"
	//打开文件，没有就创建一个，在linux 下rw可读可写
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("打开失败")
	}
	defer file.Close()
	var str = "我是追加的内容"
	n, e := file.WriteString(str)
	if e != nil {
		fmt.Println("write error", e)
		return
	}
	fmt.Println("写入的字节数是：", n)
}
