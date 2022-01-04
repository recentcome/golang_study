package main

import "fmt"

func main() {
	var str2 = [30]int{4445646, 20, 30, 50, 123, 1424, 564564, 54654646546}
	max := 0
	index := 0
	for i, v := range str2 {
		fmt.Printf("%v, %v\n", i, v)
		if v > max {
			max = v
			index = i
		}
	}
	fmt.Println("最大值为", max, "下标为", index)
}
