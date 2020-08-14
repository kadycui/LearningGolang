package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[2:5]
	s1[2] = 888

	s2 := s1[2:7]
	s2[2] = 999

	fmt.Println("s1 = ", s1, len(s1), cap(s1))
	fmt.Println("s2 = ", s2, len(s2), cap(s2))
	fmt.Println("array=", array)
}
