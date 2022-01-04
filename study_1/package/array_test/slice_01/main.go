package main

import "fmt"

func main() {
	var slice = make([]int, 10, 20)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	slice[5] = 30
	slice[9] = 100
	fmt.Println(slice)

}
