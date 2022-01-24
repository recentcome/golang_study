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
			cv.update()
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

//更新用户信息
func (cv *customerView) update() {
	fmt.Println("--------------修改客户-----------------")
	needBreak := false
	for {
		fmt.Println("请选择待修改客户的编号(-1退出)：")
		needUpdateID := -1
		fmt.Scanln(&needUpdateID)
		//如果直接输入-1，则退出
		if needUpdateID == -1 {
			return
		} else {
			//判断是否有该用户，没有该用户的id，则提示用户找不到，并在此让用户输入
			index := cv.customerService.FindById(needUpdateID)
			if index == -1 {
				fmt.Println("未找到该用户")
			} else {
				oldCustomer := cv.customerService.Customers[index]
				fmt.Printf("姓名（%v）:", oldCustomer.Name)
				newName := ""
				fmt.Scanln(&newName)
				fmt.Printf("性别（%v）:", oldCustomer.Gender)
				newGender := ""
				fmt.Scanln(&newGender)
				fmt.Printf("年龄（%v）:", oldCustomer.Age)
				newAge := 0
				fmt.Scanln(&newAge)
				fmt.Printf("电话（%v）:", oldCustomer.Phone)
				newPhone := ""
				fmt.Scanln(&newPhone)
				fmt.Printf("邮箱（%v）:", oldCustomer.Email)
				newEmail := ""
				fmt.Scanln(&newEmail)
				if newName != "" {
					oldCustomer.Name = newName
				}
				if newGender != "" {
					oldCustomer.Gender = newGender
				}
				if string(newAge) != "" {
					oldCustomer.Age = newAge
				}
				if newPhone != "" {
					oldCustomer.Phone = newPhone
				}
				if newEmail != "" {
					oldCustomer.Email = newEmail
				}
				if cv.customerService.Update(needUpdateID, oldCustomer) {
					fmt.Println("-----------------修改完成--------------")
					needBreak = true
				} else {
					fmt.Println("修改失败，请重新输入")
				}

			}
		}
		//退出循环
		if needBreak {
			break
		}
	}

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
