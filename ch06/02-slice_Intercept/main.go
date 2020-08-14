package main

import "fmt"

func main() {

	// s[low:high:max]     从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
	//s := []int{10, 20, 30, 0, 0}
	//
	//slice := s[0:3:5]
	//fmt.Println(slice)

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := array[:] // 不指定容量和长度
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s2 := array[3:] // 从下标3开始,到结尾
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[:6] // 从0开始,取6个元素,容量是10  常用
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[2:5] // array[2:5] 表示从下标为2的元素（包含该元素）开始取，到下标为5的元素(不包含该元素)结束。所以切片s5的长度是3 根据array切片的容量是10, 减去array[2:5]中的2
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))

}
