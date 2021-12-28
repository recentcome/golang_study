package main

import "fmt"

func main() {
	var i int = 1
	i++
	//++是需要单独执行的
	//var a int
	//a = i++
	//同上一样
	//if i++ > 10{
	//	fmt.Printf("ok")
	//}
	fmt.Println(i)
	fmt.Printf("123\n")

	//假如还有97天放假，问：xx个星期零xx天
	var (
		dayRelease = 97
		weekNum    int
		dayNum     int
	)
	weekNum = dayRelease / 7
	dayNum = dayRelease % 7
	fmt.Printf("还有%d个星期，零%d天", weekNum, dayNum)

	//定义变量保存华氏温度，华氏转摄氏温度为：5/9(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huaShi float32
	var sheshi float32
	huaShi = 216
	sheshi = 5.0 / 9 * (huaShi - 100)
	fmt.Printf("华氏温度%f对应的摄氏温度为%f\n", huaShi, sheshi)

	var a int = -1 << 2
	fmt.Println("a= ", a)
}
