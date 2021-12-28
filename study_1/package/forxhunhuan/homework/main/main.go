package main

import "fmt"

//1、实现判断一个整数，属于哪个范围：大于0 小于0 等于0

var num = 0

func main() {
	////1、判断一个整数属于哪个范围
	//fmt.Printf("请输入一个整数：")
	//fmt.Scanf("%d",&num)
	//if num >0{
	//	fmt.Printf("大于0")
	//} else if num < 0{
	//	fmt.Printf("小于0")
	//} else if num == 0{
	//	fmt.Printf("等于0")
	//} else{
	//	fmt.Printf("输入的有误")
	//}

	//2、判断一个年份是否为闰年
	//分析：闰年标准为: 能被4整除，不能被100整除为闰年  年份是整百的，必须是400的倍数
	//var year = 0
	//fmt.Printf("请输入年份：")
	//fmt.Scanln(&year)
	//if year%4==0 && year%100!=0 || year%400==0 {
	//	fmt.Println("年份为闰年")
	//} else{
	//	fmt.Println("年份为平年")
	//}

	//3、判断一个整数是否为水仙花数，各个位上数字的立方等于其本身，如：153 = 1^1 + 3^3 + 5^5
	//分析：如何取各个位上的数
	//个位数：num % 10
	//十位数：num % 100 /10
	//百位数：num % 1000 / 100
	//var num=0
	//fmt.Println("请输入一个整数:")
	//fmt.Scanln(&num)
	//fmt.Println((num%1000/100))
	//fmt.Println((num%100/10))
	//fmt.Println(num%10)
	//if (num%1000/100)*(num%1000/100)*(num%1000/100) +
	//	(num%100/10)*(num%100/10)*(num%100/10) +
	//	(num%10)*(num%10)*(num%10) == num{
	//	fmt.Println("是水仙花数")
	//} else{
	//	fmt.Println("不是水仙花")
	//}

	//4、保存用户名和密码，判断用户名是否为"张三",密码是否为1234，如果是提示成功，否则提示登录失败
	var usesName = "张三"
	var loginName = ""
	var passWord = "1234"
	var loginPass = ""
	fmt.Println("请输入用户名：")
	fmt.Scanln(&loginName)
	fmt.Println("请输入密码：")
	fmt.Scanln(&loginPass)
	if usesName == loginName && passWord == loginPass {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}

}
