package main

import "fmt"

func main() {
	//age := 16

	// if age := 19; age > 18 {
	// 	fmt.Println("已经成年了")
	// }else {
	// 	fmt.Println("还没有成年")
	// }

	// 多个条件判断
	//if age > 35 {
	//	fmt.Println("人到中年")
	//}else if age > 18{
	//	fmt.Println("已经成年了")
	//}else {
	//	fmt.Println("好好学习")
	//}

	var money float64

	fmt.Println("请输入公交卡的充钱数:")
	fmt.Scanf("%f\n", &money)

	if money >= 2 {
		var seatCount int
		fmt.Println("请输入空座位的数量:")
		fmt.Scanf("%d", &seatCount)

		if seatCount > 0 {
			fmt.Println("请坐!")
		} else {
			fmt.Println("请站着!")
		}
	} else {
		fmt.Println("余额不足")
	}

}
