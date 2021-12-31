package main

import (
	"fmt"
)

func main() {
	str1 := []int{12, 23, 236, 2, 124}
	//1、先把最大值放最后
	//2、先定义一个中间变量temp，str1[0]和str1[1]比较，如果str1[1]小于str1[0]，则把str1[0]的值赋给temp
	//3、str1[1] 和str1[0]做交换，然后temp存放较大的数
	temp := 0 //下标
	//下面是比较两个数的值，
	//if str1[0] > str1[1] {
	//	temp = str1[0]
	//	str1[0] = str1[1]
	//	str1[1] = temp
	//}
	//冒泡算法
	for j := 0; j < len(str1)-1; j++ { //外循环长度-1次
		for i := 0; i < len(str1)-1-j; i++ { //内循环长度-1 -外循环的次数
			if str1[i] > str1[i+1] {
				temp = str1[i]
				str1[i] = str1[i+1]
				str1[i+1] = temp
			}
		}
	}

	fmt.Println(str1)
}
