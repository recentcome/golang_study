package service

import (
	"fmt"
	"gocode/study_1/package/costumerManage/model"
)

type CustomerService struct {
	//定义客户切片，不确认有多少客户
	Customers []model.Customer
	//用户的id,可以作为有多少用户的标识
	CustomerNum int
}

func NewCustomerService() *CustomerService {
	customer := model.NewCustomer(1, "张三", "男", 28, "112", "zs@sohu.com")
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customerService.Customers = append(customerService.Customers, customer)
	return customerService
}

//查看客户详情
func (cs *CustomerService) List() []model.Customer {
	return cs.Customers
}

//添加客户，返回添加是否成功
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.CustomerNum++
	customer.Id = cs.CustomerNum
	cs.Customers = append(cs.Customers, customer)
	return true
}

//查找用户，返回对应下标，没有找到返回-1
func (cs *CustomerService) FindById(userId int) int {
	index := -1
	for i := 0; i < len(cs.Customers); i++ {
		if userId == cs.Customers[i].Id {
			//找到了
			index = i
		}
	}
	return index
}

//删除客户，返回对应下标，返回-1就会失败。
func (cs *CustomerService) Delete(userId int) bool {
	index := cs.FindById(userId)
	if index != -1 {
		cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
		return true
	}
	return false
}

//更新客户信息，功能说明，需要用户选择修改的客户编号，-1退出
//选择客户编号后，展示出该客户的信息，让用户输入对应的字段，若直接回车表示不修改
//最后提示修改是否完成。
//1、用户输入客户编号后，返回客户的详细信息。这里可以通过findbyid查询出用户，对用户的信息进行展示即可。
//2、根据用户的id，前端的输入值保存为一个新的custom，也一起传入，若不存在id，则直接返回,，
func (cs *CustomerService) Update(userId int, customer model.Customer) bool {
	index := cs.FindById(userId)
	fmt.Println("index is", index)
	if index == -1 {
		return false
	} else {
		fmt.Println(cs.Customers[index])
		fmt.Println(customer)
		if cs.Customers[index].Id == customer.Id {

		}
		cs.Customers[index] = customer
	}
	return true
}
