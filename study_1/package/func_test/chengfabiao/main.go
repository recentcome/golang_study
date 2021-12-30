package main

import (
	"fmt"
)

//分析：99乘法表 1 * 1 = 1  1*2=2 2*2=4
func chengFa(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("请输入需要打印的乘法表的个数：")
	num := 0
	fmt.Scanln(&num)
	chengFa(num)
}
