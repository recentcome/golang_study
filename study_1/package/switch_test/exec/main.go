package main

import "fmt"

func main() {
	//把小写类型的char型转换为大写
	//var str1 byte
	//fmt.Println("请输入你要转换的小写字母->>")
	//_, _ = fmt.Scanf("%c",&str1)
	//fmt.Printf("%v\n",str1)
	//switch str1 {
	//case 'a':
	//	fmt.Println("A")
	//case 'b':
	//	fmt.Println("B")
	//case 'c':
	//	fmt.Println("C")
	//case 'd':
	//	fmt.Println("D")
	//default:
	//	fmt.Println("other")
	//}
	//

	//对成绩大于60分的，输出合格，低于60分的输出不合格  输入的成绩不能大于100分
	fmt.Println("请录入学生成绩->")
	var str1 int
	_, _ = fmt.Scanf("%d", &str1)
	switch {
	case str1 > 100:
		fmt.Println("有误")
	case str1 >= 60:
		fmt.Println("及格")
	case str1 >= 0:
		fmt.Println("不及格")
	default:
		fmt.Println("有误")
	}
}
