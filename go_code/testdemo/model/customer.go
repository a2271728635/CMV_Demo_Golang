package model

import (
	"fmt"
)

/*
	Author : 秦雅楠
*/

//Customer 是一个客户信息结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//NewCustomer 是一个工厂函数，返回一个客户结构体Customer实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//NewCustomer2 是一个工厂函数，去id，由程序分配，返回一个客户结构体Customer实例
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//GetInfo 返回用户信息,格式化的字符串
func (modelthis Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", modelthis.Id, modelthis.Name,
		modelthis.Gender, modelthis.Age, modelthis.Phone, modelthis.Email)
	return info
}
