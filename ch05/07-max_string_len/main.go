package main

import "fmt"

func GetLongst(names [5]string) string {
	var max string
	max = names[0]
	for i := 1; i < len(names); i++ {
		if len(names[i]) > len(max) {
			max = names[i]
		}
	}
	return max
}

func main() {
	str := [...]string{"姚明", "迈克尔乔丹", "雷吉米勒", "蒂牡丹", "科比布莱恩特"}
	name := GetLongst(str)
	fmt.Println(name)
}
