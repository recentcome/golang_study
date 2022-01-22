package utils

import (
	"fmt"
)

type FamilyAccount struct {
	flag bool
	//定义四个变量存放账户金额、收支金额、说明 和详情
	details         string
	balance         float64
	money           float64
	note            string
	flagzhichu      bool
	transferMoney   float64
	transferAccount string
	username        string
	password        string
}

var accountList = []string{"001", "002", "003"}

var userName = "hello"
var passWord = "nihao"

//设置工厂函数，外界可以调用这个方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		flag:            true,
		details:         "收支\t账户金额\t收支金额\t说明\n",
		balance:         10000.0,
		money:           0.0,
		note:            "",
		transferMoney:   0.0,
		transferAccount: "",
		flagzhichu:      false,
		username:        "",
		password:        "",
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

func (familyAccount *FamilyAccount) TranceAccount() {
	//转账给其他用户
	//大致样子为1、请输入你要转账的账号，需要判断账号是否正确，给定几个账号
	//2、请输入你要转账的金额，需要判断金额是否比余额大
	//设置标志位
	flag := false
	for {
		fmt.Println("请输入你要转账的账号")
		fmt.Println("有这些账号可供选择", accountList)
		fmt.Scanln(&familyAccount.transferAccount)
		for _, v := range accountList {
			if v == familyAccount.transferAccount {
				flag = true
				break
			}
		}
		//退出转账循环操作
		if flag {
			break
		} else {
			fmt.Println("你输入的账号有误")
		}

	}
	for {
		fmt.Printf("请输入你要转账的金额")
		fmt.Scanln(&familyAccount.transferMoney)
		//开始转账，判断转账余额是否满足条件
		if familyAccount.transferMoney > familyAccount.balance {
			fmt.Println("余额不足，请重新输入")
		} else {
			break
		}
	}
	familyAccount.balance -= familyAccount.transferMoney
	fmt.Println("转账说明：")
	fmt.Scanln(&familyAccount.note)
	familyAccount.details += fmt.Sprintf("转账\t%v\t%v\t%v\n", familyAccount.balance,
		familyAccount.transferMoney, familyAccount.note)
	fmt.Println("转账成功")
}

func (familyAccount *FamilyAccount) Login() bool {
	fmt.Println("请输入你的账号")
	fmt.Scanln(&familyAccount.username)

	fmt.Println("请输入你的密码")
	fmt.Scanln(&familyAccount.password)

	if familyAccount.username == userName && familyAccount.password == passWord {
		fmt.Println("登录成功")
		return true
	} else {
		fmt.Println("账号或密码错误，请重新登录")
		return false
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
		fmt.Println("					   5、转   账                       ")
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
		case "5":
			familyAccount.TranceAccount()
		default:
			fmt.Println("你输入有误，请重新输入1-4")
		}
		if !familyAccount.flag {
			break
		}
	}
	fmt.Println("退出程序")
}
