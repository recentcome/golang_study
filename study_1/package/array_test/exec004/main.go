package main

import "fmt"

func main() {
	//定义一个4行4列的二维数组，逐个从键盘输入值，将第一行第四行数据进行交换，将第二行和第三行数据进行交换
	var str2 [4][4]int
	for i := 0; i < len(str2); i++ {
		for j := 0; j < len(str2[i]); j++ {
			fmt.Printf("请输入二维数组的str2的第%d行，第%d列位置的数值", i+1, j+1)
			fmt.Scanln(&str2[i][j])
		}
	}
	fmt.Println("正常数据展示为\n", str2)
	var str3 [4][4]int
	for i := 0; i < len(str2); i++ {
		str3[i] = str2[len(str2)-i-1]
	}
	fmt.Println("交换后的数据为", str3)
}
