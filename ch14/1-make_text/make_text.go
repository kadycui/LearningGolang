package main

import (
	"fmt"
	"os"
)

/*
(1)导入“os”包，创建文件，读写文件的函数都在改包
(2)指定创建的文件存放路径以及文件名。
(3)执行Create( )函数，进行文件创建
(4)关闭文件
*/

func CreateFile(path string) {
	// 创建文件返回两个值, 一个是创建的文件一个是错误信息
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	var str string
	var a int
	for i := 1; i < 10; i++ {
		str = fmt.Sprintf("i = %d\n", i+1) // 换行写入
		//n, err := f.WriteString(str)   // 写文件   第一个参数返回数据长度, 第二个我参数返回错误信息
		//fmt.Println(n)

		buf := []byte(str)

		//n, err := f.Write(buf)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(n)

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		a, err = f.WriteAt([]byte(buf), n)
		fmt.Println(a)

	}
	defer f.Close()
}

func main() {
	var filePath = "ch14/1-make_text/text/b.txt"
	CreateFile(filePath)
}
