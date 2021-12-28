package main

import "fmt"

func main() {
	//交换两个变量的值
	var a, b int
	a = 1
	b = 3
	a, b = b, a
	fmt.Printf("a=%d, b=%d", a, b)

	////求两个数的最大值
	//var num1 = 100
	//var num2 = 200
	//
	//if num1 > num2 {
	//	fmt.Printf("大的值为%d", num1)
	//} else {
	//	fmt.Printf("大的值为%d", num2)
	//}

	//求三个数的最大值 先求出两个数的最大值，再和第三个数比较
	//var numa = 3
	//var numb = 4
	//var numc = 5
	//if numa > numb && numb> numc{
	//	fmt.Println("最大的值为")
	//}
}
