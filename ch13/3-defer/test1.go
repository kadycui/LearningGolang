package main

import "fmt"

//import "fmt"
//
//func f1()(r int){
//	defer func(){
//		r++
//	}()
//	r = 0
//	return
//}
//
//func main(){
//	i := f1()
//	fmt.Println(i)
//}

func double(x int) int {
	return x + x
}

func triple(x int) (r int) {
	defer func() {
		r += x
	}()
	return double(x)
}

func main() {
	fmt.Println(triple(3)) // 9
}
