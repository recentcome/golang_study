package main

import "fmt"

//定义老猴子
type Monkey struct {
	Name string
	Age  int
}

type Flyable interface {
	Flying()
}

type Swam interface {
	Swamming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "会爬树")
}

type LittleMonkey struct {
	Monkey
}

func (little *LittleMonkey) Flying() {
	fmt.Println("通过学习", little.Name, "会飞了")
}

func (little *LittleMonkey) Swamming() {
	fmt.Println("通过学习", little.Name, "会游泳了")
}

func main() {
	money := LittleMonkey{Monkey{
		Name: "悟空",
	}}
	money.climbing()
	money.Flying()
	money.Swamming()
}
