package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Printf("res is %d", res)
}

func main() {
	test()
	fmt.Printf("main函数继续执行")
}
