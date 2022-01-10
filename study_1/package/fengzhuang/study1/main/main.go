package main

import (
	"fmt"
	"gocode/study_1/package/fengzhuang/study1/model"
)

func main() {
	p := model.NewPerson("mary")
	p.SetAge(-1)
	fmt.Println(p.GetAge())
}
