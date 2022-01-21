package main

import "fmt"

func main() {
	flag := true
	//定义四个变量存放账户金额、收支金额、说明 和详情
	details := "收支\t账户金额\t收支金额\t说明\n"
	balance := 10000.0
	money := 0.0
	note := ""
	flagzhichu := false
	for {
		var choose string
		fmt.Println("\n----------------------家庭收支记账软件----------------------")
		fmt.Println()
		fmt.Println("                       1、收支明细                       ")
		fmt.Println("					   2、登记收入                       ")
		fmt.Println("					   3、登记支出                       ")
		fmt.Println("					   4、退   出                       ")
		fmt.Println()
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&choose)

		switch choose {
		case "1":
			fmt.Println("------------------当前收支明细记录----------------------")
			if flagzhichu {
				fmt.Printf(details)
			} else {
				fmt.Println("当前没有收支，请来一笔吧----")
			}

		case "2":
			fmt.Println("登记收入")
			fmt.Println("收入的金额：")
			fmt.Scanln(&money)
			fmt.Println("金额说明：")
			fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("收入\t%v\t%v\t%v\n", balance, money, note)
			flagzhichu = true
		case "3":
			fmt.Println("登记支出")
			fmt.Println("支出的金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("金额说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("支出\t%v\t%v\t%v\n", balance, money, note)
			flagzhichu = true
		case "4":
			//用户确认是否退出
			exitC := ""
			fmt.Println("请确认是否退出：y/n")
			for {
				fmt.Scanln(&exitC)
				if exitC == "y" || exitC == "n" {
					break
				}
				fmt.Println("你输入有误，请重新输入y/n")
			}
			if exitC == "y" {
				flag = false
			}

		default:
			fmt.Println("你输入有误，请重新输入1-4")
		}
		if !flag {
			break
		}
	}
	fmt.Println("退出程序")

}
