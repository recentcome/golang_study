package main

//编写判断大鱼还是筛网从1990年1月1日起开始执行
//分析：1、窗口需要用户手动输入年月日，fmt.Scanf接收3个变量年月日，然后修改为时间函数传入判断大鱼还是筛网的函数中。
// 	 	2、判断的函数isFishOrNet计算出传入的时间戳，和1990年1月1日的时间戳 相差多少天，再对5求模，
//      3、求的模后，如何判断是大鱼还是筛网。。。。
import (
	"fmt"
	"strconv"
)

//func isFishOrNet(time time.Time) string{
//	time.Unix()
//}

func main() {
	var (
		year,
		month,
		day int
	)
	fmt.Printf("请输入年")
	fmt.Scanln(&year)
	fmt.Printf("请输入月")
	fmt.Scanln(&month)
	fmt.Printf("请输入日")
	fmt.Scanln(&day)
	fmt.Printf("你输入的日期为%v-%v-%v\n", year, month, day)
	str1 := strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(day)
	fmt.Printf(str1)
}
