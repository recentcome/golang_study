package main

import "fmt"

type methods struct {
	chang int
	gao   int
}

func (cd methods) MethodUtils(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()

	}
}

func main() {
	m1 := methods{10, 20}
	m1.PrintJu()
}
