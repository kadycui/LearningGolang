package main

import "fmt"

func main() {
	var m map[int]string
	fmt.Println(len(m), m)

	m2 := make(map[int]string)
	fmt.Println(len(m2), m2)

	// 指定容量
	m3 := make(map[int]string, 3)
	fmt.Println(len(m3), m3)

	m3[1] = "张三"
	m3[2] = "李四"
	m3[3] = "王五"
	m3[4] = "赵柳"
	m3[5] = "牵起"

	fmt.Println(len(m3), m3)

	// 定义并初始化
	m4 := map[int]string{1: "牛奶", 2: "果汁", 3: "咖啡", 4: "绿茶"}
	fmt.Println(len(m4), m4)

	// 打印字典的值
	for key, value := range m4 {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}

	va, ok := m4[1]
	if ok == true {
		fmt.Println("m4[1]", va)
	} else {
		fmt.Println("key值不存在")
	}

	// 删除字典的值
	m5 := map[int]string{1: "go", 2: "python", 3: "java", 4: "c++"}
	delete(m5, 1)
	fmt.Println(m5)

	// map作为函数参数是引用传递
	m6 := map[int]string{1: "golang", 2: "python", 3: "java", 4: "c++"}
	DeleMap(m6)
	fmt.Println(m6)

}

func DeleMap(m map[int]string) {
	delete(m, 2)
}
