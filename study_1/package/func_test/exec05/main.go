package main

import (
	"fmt"
	"os"
)

func computer(num int) {
	a := 10
	b := 5
	switch num {
	case 1:
		fmt.Printf("%v+%v=%v\n", a, b, a+b)
	case 2:
		fmt.Printf("%v-%v=%v\n", a, b, a-b)
	case 3:
		fmt.Printf("%v*%v=%v\n", a, b, a*b)
	case 4:
		fmt.Printf("%v/%v=%v\n", a, b, a/b)
	case 0:
		fmt.Printf("程序退出！")
		os.Exit(0)
	default:
		fmt.Printf("你输入的数有误\n")
	}

}

//打印小小计算器 只支持加减乘除
func main() {
	for {
		fmt.Println("————————小小计算器————————")
		fmt.Println("1.加法")
		fmt.Println("2.减法")
		fmt.Println("3.乘法")
		fmt.Println("4.除法")
		fmt.Println("0.退出")
		fmt.Println("请选择：")
		num := 0
		fmt.Scanln(&num)
		computer(num)
	}

}
