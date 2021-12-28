package main

import "fmt"

//求斐波那契数  1,1,2,3,5,8,13...
//分析：
// 1、位置对应的是自增个数（对应给定的个数）
// 2、前两个数的和等于第三个数  大于3时，第三个数等于第一个数+第二个数
// 3、先搞定3个数的方法
func diGui(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return diGui(n-1) + diGui(n-2)
	}
}

//2、求函数值 一直f(1) = 3 ; f(n) = 2* f(n-1)+1

func hanSshu(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*hanSshu(n-1) + 1
	}
}

//3、猴子吃桃子问题
//一堆桃子，第一天吃一半，并多吃一个，以后每天都吃一般，多吃一个，第十天再次时，只有一个桃子，问最初多少个桃子
//分析：假设有N个桃子，第一天吃N/2+1个桃子，第二天吃（N-N/2+1）/2+1个桃子 第十天时，只有一个桃子
//反推，第10天吃1个桃子  第9天吃 （第10天的桃子+1）*2  第8天吃 (第9天吃的桃子+1) *2
// 第9天有4个桃子，吃一半	+1 还剩1个桃子
//第8天有个10桃子，吃一半	+1 还剩4个桃子
//第7天有个22个桃子，吃一半		+1 还剩10个桃子
//第n天的桃子数据 peach(n) = (peach)
func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
func main() {

	fmt.Println(diGui(6))
	fmt.Println(hanSshu(7))
	fmt.Println(peach(10))
}
