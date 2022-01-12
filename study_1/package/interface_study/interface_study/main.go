package main

import (
	"fmt"
)

type AInterface interface {
	test01()
	test02()
}

type BInterface interface {
	test01()
	test03()
}

type CInterface interface {
	AInterface
	BInterface
}

type Student struct {
}

func (stu *Student) test01() {
}
func (stu *Student) test02() {
}
func (stu *Student) test03() {
}
func main() {
	var stu Student
	var a AInterface = &stu
	fmt.Println(a)

	var c CInterface = &stu
	fmt.Println(c)
}
