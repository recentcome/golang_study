package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println("123")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	test1 := make([]int, 10)
	for i := range test1 {
		test1[i] = 1 << uint(i) // ==2**i
	}

	for _, value := range test1 {
		fmt.Printf("%d\n", value)
	}
}
