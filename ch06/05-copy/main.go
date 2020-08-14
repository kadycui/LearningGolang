package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}

	// 将第二个切片里面的元素，拷贝到第一个切片中。
	copy(dstSlice, srcSlice)
	fmt.Println("dst=", dstSlice)
	copy(srcSlice, dstSlice)
	fmt.Println("src=", srcSlice)
}
