package main

import (
	"fmt"
	"go_code/testdemo/model"
	"go_code/testdemo/service"
)

/*
	Author : 廖怡
*/

type date struct {
	//定义必要字段
	falg bool   //是否应该循环显示菜单
	key  string //接收用户输入信息
	//添加一个字段customerservice
	customerService *service.CustomerService
}

//显示所有的客户信息
func (viewthis *date) list() {
	//首先，获取当前所有的客户信息(在切片中)
	customers := viewthis.customerService.List()
	//显示

	fmt.Println("-----------客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------客户列表完成-----------")
}

//根据用户输入的信息 添加新客户
func (viewthis *date) add() {
	fmt.Println("-----------添加客户-----------")

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer结构体
	//注意：id号为自增，不能由用户输入
	customer2 := model.NewCustomer2(name, gender, age, phone, email)
	if viewthis.customerService.Add(customer2) {
		fmt.Println("-----------添加完成-----------")
	} else {
		fmt.Println("-----------添加失败-----------")
	}

}

//根据输入的客户编号，删除客户
func (viewthis *date) del() {
	id := -1
	fmt.Println("-----------删除客户-----------")
	fmt.Println("输入你想删除的客户编号(输入-1为退出删除菜单)")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("你已退出")
		return
	}
	for {
		viewthis.key = ""
		fmt.Println("请确认是否删除(y/n)")
		fmt.Scanln(&viewthis.key)
		if viewthis.key == "y" {
			if viewthis.customerService.Del(id) {
				fmt.Printf("已删除编号为%d的客户 \n", id)
				break
			} else {
				fmt.Println("删除失败")
				break
			}
		} else if viewthis.key == "n" {
			fmt.Println("已取消删除")
			break
		}
	}

}

//根据输入的客户编号，编辑客户
func (viewthis *date) edit() {
	id := -1
	fmt.Println("-----------编辑客户-----------")
	fmt.Println("输入你想删除的客户编号(输入-1为退出编辑菜单)")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("你已退出")
		return
	}
	if viewthis.customerService.FindById(id) == -1 {
		fmt.Println("没有此用户，编辑失败")
	} else {
		fmt.Println("姓名:")
		name := ""
		fmt.Scanln(&name)

		fmt.Println("性别:")
		gender := ""
		fmt.Scanln(&gender)

		fmt.Println("年龄:")
		age := 0
		fmt.Scanln(&age)

		fmt.Println("电话:")
		phone := ""
		fmt.Scanln(&phone)

		fmt.Println("邮箱:")
		email := ""
		fmt.Scanln(&email)
		customer2 := model.NewCustomer(id, name, gender, age, phone, email)
		if viewthis.customerService.Edit(customer2) {
			fmt.Println("-----------编辑完成-----------")
		}
	}
}

//exit 是否退出软件确认操作
func (viewthis *date) exit() {
	for {
		viewthis.key = ""
		fmt.Println("请确认是否退出(y/n)")
		fmt.Scanln(&viewthis.key)
		if viewthis.key == "y" {
			viewthis.falg = false
			break
		} else if viewthis.key == "n" {
			fmt.Println("已取消 【退出】")
			viewthis.falg = true
			break
		}
	}
}

//menuview 是view层的主菜单显示
func (viewthis *date) menuview() {
	for {
		fmt.Println("-----------客户管理系统-----------")
		fmt.Println("            1.添加客户")
		fmt.Println("            2.修改客户")
		fmt.Println("            3.删除客户")
		fmt.Println("            4.客户列表")
		fmt.Println("            5.退出")
		fmt.Println("请输入1-5")
		fmt.Scanln(&viewthis.key)
		switch viewthis.key {
		case "1":
			fmt.Println("添   加   客   户")
			viewthis.add()
		case "2":
			fmt.Println("修   改   客   户")
			viewthis.edit()
		case "3":
			fmt.Println("删   除   客   户")
			viewthis.del()
		case "4":
			fmt.Println("客   户   列   表")
			viewthis.list()
		case "5":
			viewthis.exit()
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if viewthis.falg == false {
			fmt.Println("你已退出")
			break
		}
	}
}

func main() {
	fmt.Println("软件已运行__~~~~")
	//创建实例并运行显示主菜单
	Date := date{
		falg: true,
		key:  "",
	}
	//这里完成对Date结构体的customerService字段的初始化
	Date.customerService = service.NewCustomerService()

	(&Date).menuview()
}
