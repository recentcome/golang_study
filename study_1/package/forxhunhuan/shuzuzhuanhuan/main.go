package main

import "fmt"

func main() {
	//编写数组 1 2 3
	//         4 5 6
	//        7 8 9
	//先打印1-3   再打印
	//values :=[3][3]int
	//for i:=1;i<=3 ;i++  {	//层数
	//	fmt.Printf("%v ",i)
	//}
	//fmt.Println("")
	//for i:=4; i<=6;i++ {
	//	fmt.Printf("%v ",i)
	//}
	//fmt.Println("")
	//for i:=7; i<=9;i++ {
	//	fmt.Printf("%v ",i)
	//}
	//fmt.Println("")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%v ", i)
		if i%3 == 0 {
			fmt.Println("")
		}
	}
}
