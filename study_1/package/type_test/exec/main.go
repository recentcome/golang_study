package main

import "fmt"

//可以自定义数据类型，将int型定义为自己的名字myInt

type myInt int

var num myInt = 0

//定义函数的名字，简化
type myFunc func(int, int) int

func setFunc(funvar myFunc, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	num = 1
	fmt.Println("num=", num)
	sumFunc := setFunc(getSum, 2, 4)
	fmt.Println("result=", sumFunc)

}

func getSum(i int, i2 int) int {
	return i + i2
}
