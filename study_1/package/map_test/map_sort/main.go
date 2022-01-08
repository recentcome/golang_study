package main

import "fmt"

func main() {
	map1 := make(map[int]int)
	map1[10] = 100
	map1[3] = 12
	map1[34] = 24
	map1[213] = 445
	fmt.Println(map1)
}
