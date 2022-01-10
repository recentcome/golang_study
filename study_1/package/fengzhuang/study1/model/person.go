package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  int
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄设置错误")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal int) {
	if sal > 3000 && sal < 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水设置错误")
	}
}

func (p *person) GetSal() int {
	return p.sal
}
