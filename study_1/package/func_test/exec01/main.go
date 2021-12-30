package main

import "fmt"

//循环打印输入的月份的天数 【使用continue实现】
//输入年、月，输入
func main() {
	for {
		yearNum := 0
		monthNum := 0
		days := 0
		fmt.Printf("请输入年：")
		fmt.Scanln(&yearNum)
		fmt.Printf("请输入月")
		fmt.Scanln(&monthNum)
		//判断是否为闰年
		switch monthNum {
		case 1, 3, 5, 7, 8, 10, 12:
			days = 31
		case 2, 4, 6, 9, 11:
			days = 30
		default:
			fmt.Printf("请输入正确的月份\n")
			continue
		}

	}
}
