package main

import (
	"fmt"
)

func main() {
	fmt.Println("你好")
	//rec1是匿名函数，只能调用一次
	rec1 := func(n1, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println(rec1)

	//rec2可以调用多次
	rec2 := func(n2, n3 int) int {
		return n2 + n3
	}
	a := rec2(2, 3)
	b := rec2(44, 11)
	fmt.Println("a", a, b)
}
