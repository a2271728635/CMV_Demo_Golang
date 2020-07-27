package service

import (
	"go_code/testdemo/model"
)

/*
	Author : 秦雅楠
*/

//CustomerService 完成对Customer的操作，包括增删改查
type CustomerService struct {
	customer []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//NewCustomerService 实例化model中的Customer，并添加一个初始用户
//返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到客户在 customer []model.Customer 切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "135", "135@qq.com")
	customerService.customer = append(customerService.customer, customer)
	return customerService
}

//List 返回客户切片
func (servicethis *CustomerService) List() []model.Customer {
	return servicethis.customer
}

//Add 添加客户到customer切片
func (servicethis *CustomerService) Add(customer model.Customer) bool {
	servicethis.customerNum++
	customer.Id = servicethis.customerNum
	servicethis.customer = append(servicethis.customer, customer)
	return true
}

//FindById 根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (servicethis *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(servicethis.customer); i++ {
		if servicethis.customer[i].Id == id {
			index = i
		}
	}
	return index
}

//Del 从customer切片中删除客户，根据传入编号
func (servicethis *CustomerService) Del(id int) bool {
	index := servicethis.FindById(id)
	//如果为-1，说明没有这个客户
	if index == -1 {
		return false
	}
	servicethis.customer = append(servicethis.customer[:index], servicethis.customer[index+1:]...)
	return true
}

//Edit 从customer切片中编辑客户，根据传入的新信息
func (servicethis *CustomerService) Edit(customer model.Customer) bool {
	index := servicethis.FindById(customer.Id)
	//如果为-1，说明没有这个客户
	if index == -1 {
		return false
	}

	servicethis.customer[index] = customer

	return true
}
