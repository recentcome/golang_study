package main

import "fmt"

func main() {
	var (
		name string
		//age byte
		//sal float32
		//isPass bool
	)
	fmt.Println("请输入年龄")
	fmt.Scanln(&name)

	fmt.Println("年龄是", name)
}
