package main

import "fmt"

func main() {
	////打印1-100之间所有是9的倍数的整数的个数及总和
	//var numI =0
	//var numSum = 0
	//for i:=1; i<=100;i++  {
	//	if i%9 == 0{
	//		fmt.Println(i)
	//		numSum +=i
	//		numI++
	//	}
	//}
	//fmt.Printf("1-100中9的倍数的整数有%d个\n", numI)
	//fmt.Printf("1-100中9的倍数的整数的总和是%d", numSum)

	//完成下面的表达式输出
	//0 + 6 = 6
	//1 + 5 = 6
	//2 + 4 = 6
	//3 + 3 = 6
	//4 + 2 = 6
	//5 + 1 = 6
	var n = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v \n", i, n-i, n)

	}

}
