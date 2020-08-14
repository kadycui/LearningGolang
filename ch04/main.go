package main

// 导包
//  GOPATH=C:\GoWork
// C:\GoWork\src\LearningGoLang

import (
	"LearningGoLang/ch04/lib"
	"LearningGoLang/ch04/product"
	"LearningGoLang/ch04/userinfo"
	"fmt"
)

func main() {
	fmt.Println("开始导包")
	userinfo.AddUser()
	product.AddProduct()
	lib.Test1()
	lib.Test2()

}
