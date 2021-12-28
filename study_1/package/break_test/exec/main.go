package main

import "fmt"

func main() {
	//1、100以内求和，求出当和第一次大于20的当前数
	//var sum=0
	//for i:=1;i<=100;i++ {
	//	sum+=i
	//	if sum> 20{
	//		fmt.Println("和为",sum)
	//		fmt.Println("当前的数为", i )
	//		break
	//	}
	//}

	//2、实现登录验证，三次机会，如果用户名为“张无忌”，密码“888” 提示登录成功，否则提示还有几次机会
	var usrName = "张无忌"
	var passWord = "8888"
	var loginName = ""
	var loginPass = ""
	for i := 3; i >= 0; i-- {
		fmt.Println("请输入用户名")
		fmt.Scanln(&loginName)
		fmt.Println("请输入密码")
		fmt.Scanln(&loginPass)
		if usrName == loginName && passWord == loginPass {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("你还有%d次机会", i)
		}
	}

}
