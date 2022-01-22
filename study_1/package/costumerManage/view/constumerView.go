package main

import (
	"fmt"
	"gocode/study_1/package/costumerManage/model"
	"gocode/study_1/package/costumerManage/service"
)

type customerView struct {
	choose          string
	loop            bool
	customerService *service.CustomerService
}

func (cv *customerView) mainMenu() {
	cv.loop = false
	for {
		fmt.Println("\n-----------客户信息管理软件-------------")
		fmt.Println("          1、添加客户")
		fmt.Println("          2、修改客户")
		fmt.Println("          3、删除客户")
		fmt.Println("          4、客户列表")
		fmt.Println("          5、退   出")
		fmt.Println("          请选择（1-5）：")

		fmt.Scanln(&cv.choose)
		switch cv.choose {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		}
		if cv.loop {
			break
		}
	}

}

//显示所有的客户信息
func (cv *customerView) list() {
	customers := cv.customerService.List()
	//显示
	fmt.Println("--------------客户列表------------------")
	//输出客户信息
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------客户列表完成------------------")
}

//显示所有的客户信息
func (cv *customerView) add() {
	fmt.Println("---------------添加客户---------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("---------------添加成功---------------")
	}
}

func (cv *customerView) delete() {
	fmt.Println("---------------删除客户---------------")
	//定义删除操作的标志位
	key := false
	for {
		fmt.Println("请选择待删除客户编号（-1退出）：")
		num := -1
		fmt.Scanln(&num)
		if num == -1 {
			return
		}
		fmt.Println("确认是否删除（Y/N）：")
		for {
			isY := ""
			fmt.Scanln(&isY)
			if isY == "y" || isY == "Y" || isY == "N" || isY == "n" {
				if isY == "Y" || isY == "y" {
					if cv.customerService.Delete(num) {
						key = true
					} else {
						fmt.Println("你要删除的客户不存在")
					}
					break
				} else {
					return
				}
			} else {
				fmt.Println("你输入错误，确认是否删除（Y/N）：")
			}
		}

		if key {
			fmt.Println("---------------删除完成---------------")
			break
		}

	}
}

func (cv *customerView) exit() {
	for {
		fmt.Println("请确认是否退出Y/N：")
		for {
			isY := ""
			fmt.Scanln(&isY)
			if isY == "y" || isY == "Y" || isY == "N" || isY == "n" {
				if isY == "Y" || isY == "y" {
					cv.loop = true
					break
				} else {
					return
				}
			} else {
				fmt.Println("你输入错误，确认是否退出（Y/N）：")
			}
		}
		if cv.loop {
			break
		}
	}
}

//待完成
func (cv *customerView) changeUser() {

}
func main() {
	fmt.Println("客户管理系统开始")
	mM := customerView{
		choose: "",
		loop:   false,
	}
	mM.customerService = service.NewCustomerService()
	mM.mainMenu()

}
