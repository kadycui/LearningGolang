package main

import "fmt"

func main() {
	s := make([]int, 5, 8)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	fmt.Println(s)
	s[0] = 1
	s = append(s, 2, 3, 4, 5) // 超出容量后两倍扩容
	fmt.Println(s)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
