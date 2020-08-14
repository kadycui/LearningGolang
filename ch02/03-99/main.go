package main

import "fmt"

func main() {
	//var i int
	//var j int
	//for i = 1; i <= 9; i++ {
	//	for j=1; j<=i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println("")
	//	/*
	//	*
	//	**
	//	***
	//	****
	//	*****
	//	******
	//	*******
	//	********
	//	*********
	//	*/
	//}

	var i int
	var j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println(" ")
	}
}
