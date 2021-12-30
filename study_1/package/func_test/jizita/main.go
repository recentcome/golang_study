package main

import "fmt"

//终端输入一个整数，打印对应的金字塔
//1、打印矩形 Num传入的是层数
//2、打印对应的个数 1-》1  2——》3 3——》5 4——》7  2num-1
//3、打印空格的个数  空行和层数对应关系  1-> 6个空格  2-> 5个空格 。。。7-> 无空格
func createCute(num int) {
	for i := 1; i <= num; i++ { //对应层数 需要打多少层
		for k := num - 1; k >= i; k-- { //第一次进来需要打6个空格，
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ { //对应每一行的个数
			//判断一下层数和空行的关系，如果
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("请输入需要打印的金字塔高度：")
	num := 0
	fmt.Scanln(&num)
	createCute(num)
}
