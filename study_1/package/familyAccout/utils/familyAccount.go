package utils

import "fmt"

type FamilyAccount struct {
	flag bool
	//定义四个变量存放账户金额、收支金额、说明 和详情
	details    string
	balance    float64
	money      float64
	note       string
	flagzhichu bool
}

//设置工厂函数，外界可以调用这个方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		flag:       true,
		details:    "收支\t账户金额\t收支金额\t说明\n",
		balance:    10000.0,
		money:      0.0,
		note:       "",
		flagzhichu: false,
	}
}

func (familyAccount *FamilyAccount) showDetail() {
	fmt.Println("------------------当前收支明细记录----------------------")
	if familyAccount.flagzhichu {
		fmt.Printf(familyAccount.details)
	} else {
		fmt.Println("当前没有收支，请来一笔吧----")
	}
}

func (familyAccount *FamilyAccount) income() {
	fmt.Println("登记收入")
	fmt.Println("收入的金额：")
	fmt.Scanln(&familyAccount.money)
	fmt.Println("金额说明：")
	fmt.Scanln(&familyAccount.note)
	familyAccount.balance += familyAccount.money
	familyAccount.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", familyAccount.balance,
		familyAccount.money, familyAccount.note)
	familyAccount.flagzhichu = true
}

func (familyAccount *FamilyAccount) outcome() {
	fmt.Println("登记支出")
	fmt.Println("支出的金额：")
	fmt.Scanln(&familyAccount.money)

	if familyAccount.money > familyAccount.balance {
		fmt.Println("余额不足")
		return
	}
	familyAccount.balance -= familyAccount.money
	fmt.Println("金额说明：")
	fmt.Scanln(&familyAccount.note)
	familyAccount.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", familyAccount.balance,
		familyAccount.money, familyAccount.note)
	familyAccount.flagzhichu = true
}

func (familyAccount *FamilyAccount) exit() {
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
		familyAccount.flag = false
	}
}

func (familyAccount *FamilyAccount) MainMenu() {
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
			familyAccount.showDetail()

		case "2":
			familyAccount.income()
		case "3":
			familyAccount.outcome()
		case "4":
			familyAccount.exit()

		default:
			fmt.Println("你输入有误，请重新输入1-4")
		}
		if !familyAccount.flag {
			break
		}
	}
	fmt.Println("退出程序")
}
