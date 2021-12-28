package main

import "fmt"

func main() {
	////1、打印1--100之内的奇数
	//for i:=1;i<=100;i++{
	//	if i%2 ==0 {
	//		continue
	//	} else {
	//		fmt.Println("当前的数为",i )
	//	}
	//}

	//2、键盘读入个数不定的整数，并判断读入的整数和负数的个数，输入为0时结束程序
	var positive_num = 0
	var negative_num = 0
	for {
		var num = 0
		fmt.Println("请输入需要录入的整数")
		fmt.Scanln(&num)
		if num > 0 {
			positive_num += 1
			continue
		} else if num < 0 {
			negative_num += 1
			continue
		} else if num == 0 {
			break
		} else {
			fmt.Println("你输入的数据不正确")
			continue
		}
	}
	fmt.Println("正数的个数为", positive_num)
	fmt.Println("负数的个数为", negative_num)

	fmt.Println("OK1")
	fmt.Println("OK2")
	goto label //跳过
	fmt.Println("OK3")
	fmt.Println("OK4")
	fmt.Println("OK5")
label:
	fmt.Println("OK6")

}
