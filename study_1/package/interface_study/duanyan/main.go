package main

import "fmt"

func main() {
	var a interface{}
	b := float32(24)
	a = b
	fmt.Println(a)
	c, ok := a.(float64)
	if ok {
		fmt.Printf("转化后的类型为%T,值为%v", c, c)
	} else {
		fmt.Println("转化失败")
	}
	fmt.Println("继续执行")
}
