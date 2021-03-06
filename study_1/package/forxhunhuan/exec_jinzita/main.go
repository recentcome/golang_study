package main

import "fmt"

func main() {
	//使用for循环打印出金字塔
	//编程思路
	//1、打印一个矩形
	/*
	***
	***
	***
	 */
	//for i:=1;i<=3 ;i++  {	//层数
	//	for j:=1;j<=3 ;j++ { //打印的个数
	//		fmt.Printf("*")
	//	}
	//	fmt.Println("")	//打印换行
	//}
	//}

	//2、打印半个金字塔
	/*
	*		第一行打印1个星
	**		第二行打印2个星
	***		第三行打印3个星			规律：第N行打印N个星
	 */
	//
	//for i:=1;i<=3 ;i++  {	//层数
	//	for j:=1;j<=i ;j++ { //打印的个数  //第N行打印N个星
	//		fmt.Printf("*")
	//	}
	//	fmt.Println("")	//打印换行
	//	}
	//}

	//3、打印整个金字塔
	/*
		   *		第一行打1个星，	第一行打印3个空格
		  ***    第二行打三个星，第二行打印2个空格
		 *****   第三行打5个星， 第三行1个空格
			规律：
			1、第N行打2N-1个星
			2、1 - > 2
	*/
	//定义层数
	var num = 10
	for i := 1; i <= num; i++ { //层数
		////打印空行和层数相反
		for k := num; k >= i; k-- {
			fmt.Printf(" ")
		}
		//分析：层数和空格数的关系为第N层对应的空格数为层数
		fmt.Printf(" ")
		for j := 1; j <= 2*i-1; j++ { //打印的个数  //第N行打印N个星
			if j == 1 || j == 2*i-1 || i == num {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
			//if j==i-j || j==2*i-1 {
			//	fmt.Printf("*")
			//} else {
			//	fmt.Printf(" ")
			//}
		}
		fmt.Println("") //打印换行
	}
}

//4、打印空心金字塔
/*
   *
  * *		//每行的第一个和最后一个打星
 *****
*/
