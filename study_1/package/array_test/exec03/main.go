package main

import "fmt"

//求一个数组的和和平均值
func main() {
	sum := 0
	var str1 = [20]int{1, 2, 3, 4, 5, 6, 8, 90, 123, 2145, 1233}
	for _, v := range str1 {
		sum += v
	}
	fmt.Printf("和为%d", sum)
	fmt.Printf("平均值为%d", sum/len(str1))
}
