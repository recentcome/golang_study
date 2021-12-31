package main

import "fmt"

func fbn(n int) []int {
	var str1 = make([]int, n)
	str1[0] = 1
	str1[1] = 1
	for i := 2; i < n; i++ {
		str1[i] = str1[i-1] + str1[i-2]
	}
	return str1
}

func main() {
	fmt.Println(fbn(10))
}
