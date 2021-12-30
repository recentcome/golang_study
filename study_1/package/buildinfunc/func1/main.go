package main

import "fmt"

func main() {
	var i = 100
	fmt.Printf("i的值为%v,i的地址为%v,i的类型为%T\n", i, &i, i)

	var j = new(int)
	fmt.Printf("j的值为%v,j的地址为%v,%T\n", j, &j, j)
	*j = 100
	fmt.Printf("j指针指向的值为%v", *j)
}
