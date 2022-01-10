package main

import (
	"fmt"
	"gocode/study_1/package/fengzhuang/exec01/model"
)

func main() {
	ac := model.Account{Account: "nihao"}
	ac.SetAccount("gs2001")
	ac.SetPwd("666666")
	ac.SetSal(300)

	fmt.Println(ac.GetAccount())
	fmt.Println(ac.GetPwd())
	fmt.Println(ac.GetSal())
}
