package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机猜数游戏
	//随机生成一个1-100的整数  有10次机会
	//i :=rand.Intn(100)
	//种子生成，使用系统时间的不确定性来初始化
	rand.Seed(time.Now().Unix())
	i := rand.Intn(100)
	num := 0
	fmt.Printf("%T %T", i, num)
	for j := 1; j <= 10; j++ { //猜数的次数 只有10次
		fmt.Printf("第%v次,请输入你要猜的数", j)
		fmt.Scanln(&num)
		if num == i {
			switch j {
			case 1:
				fmt.Println("你真是个天才")
				break
			case 2, 3:
				fmt.Println("你很聪明，赶上我了")
				break
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("一般般")
				break
			default:
				fmt.Printf("说你点啥好呢")
			}
			break
		} else if num > i {
			fmt.Printf("大了")
		} else if num < i {
			fmt.Printf("小了")
		} else {
			fmt.Printf("你输入的值不正确，请输入1-100整数")
		}

	}
	//fmt.Println("生成的随机数为",rand.Intn(100))
}
