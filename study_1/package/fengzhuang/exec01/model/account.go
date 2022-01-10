package model

import "fmt"

type Account struct {
	Account string
	pwd     string
	balance float64
}

func (p *Account) SetAccount(account string) {
	lenthAccount := len(account)
	if lenthAccount >= 6 && lenthAccount <= 20 {
		p.Account = account
	} else {
		fmt.Println("账号必须在6-20位之间")
	}
}

func (p *Account) GetAccount() string {
	return p.Account
}

func (p *Account) SetPwd(pwd string) {
	lenthPwd := len(pwd)
	if lenthPwd == 6 {
		p.pwd = pwd
	} else {
		fmt.Println("密码必须是6位")
	}
}

func (p *Account) GetPwd() string {
	return p.pwd
}

func (p *Account) SetSal(balance float64) {
	if balance > 20 {
		p.balance = balance
	} else {
		fmt.Println("余额必须大于30")
	}
}

func (p *Account) GetSal() float64 {
	return p.balance
}
