package main

import "fmt"

//将斐波那契额数列放进切片中
// 1,1,3,5,8,13
func feiBo(n int) int {
	//var str []int
	if n < 3 {
		return 1
	} else {
		return feiBo(n-1) + feiBo(n-2)
	}
}

func main() {
	//编写一个函数fbn
	var str []int
	n := 8
	for i := 1; i <= n; i++ {
		str = append(str, feiBo(i))
	}
	fmt.Println(str)

}
