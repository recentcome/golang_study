package main

import (
	"fmt"
	"gocode/study_1/package/familyAccout/utils"
)

func main() {
	fmt.Println("这是一个面向对象的实现")
	flag := false
	for {
		islogin := utils.NewFamilyAccount().Login()
		if islogin {
			utils.NewFamilyAccount().MainMenu()
			flag = true
		}
		if flag {
			break
		}
	}

}
