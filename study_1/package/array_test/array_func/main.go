package main

import "fmt"

func main() {
	var array1 [10]int
	array1[0] = 10
	array1[1] = 9
	array1[2] = 8
	array1[3] = 7
	array1[4] = 6

	for i, v := range array1 {
		fmt.Println("index is :", i)
		fmt.Println("value is ", v)
	}
}
