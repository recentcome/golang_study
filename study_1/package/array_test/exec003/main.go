package main

import "fmt"

//定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
/*
	[1 2 3 4]
	[1 2 3 4]
	[1 2 3 4]
*/
//分析：最外层打印0
//只有中间的两个数为原值  arr[1][1]和arr[1][2]是原值 其余都置零
func main() {
	var str2 [3][4]int
	for i := 0; i < len(str2); i++ {
		for j := 0; j < len(str2[i]); j++ {
			fmt.Printf("请输入二维数组的str2的第%d行，第%d列位置的数值", i+1, j+1)
			fmt.Scanln(&str2[i][j])
		}
	}
	fmt.Println("正常数据展示为\n", str2)
	// 当前打印是这样的
	// [[1 2 3 4] [1 2 3 4] [1 2 3 4]]
	for i := 0; i < len(str2); i++ {
		//fmt.Println(str2[i])
		for j := 0; j < len(str2[i]); j++ {
			if i == 1 && (j == 1 || j == 2) {
				fmt.Printf("%d ", str2[i][j])
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println("")
	}

}
