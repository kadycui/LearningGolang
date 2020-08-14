package main

import "fmt"

func main() {
	var sum int

	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
