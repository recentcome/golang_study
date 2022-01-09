package main

import "fmt"

//结构体，编程一个方法，在方法中打印10*8的矩形
type MethodUtils struct {
	chang int
	gao   int
}

func (m MethodUtils) PrintJu() {
	for i := 0; i < m.chang; i++ {
		for j := 0; j < m.gao; j++ {
			fmt.Printf("*")
		}
		fmt.Println()

	}
}
func main() {
	m1 := MethodUtils{10, 8}
	m1.PrintJu()
}
