package main

import (
	"awesomeProject/customerProject426/service"
	"fmt"
)

type customerView struct {
	//定义必要字段
	key  string
	loop bool
	//customerService字段
	customerService *service.CustomerService
}

//显示客户列表
func (v *customerView) list() {
	//首先获取所有的客户信息（切片）
	customers := v.customerService.List()
	//显示
	fmt.Println("--------------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	//for i := 0; i < len(customers); i++ {
	//	fmt.Println(customers[i].GetInfo())
	//}
	for _, customer := range customers {
		customer.GetInfo()
	}
	fmt.Println("--------------------客户列表完成-------------")
}

//添加客户
func (v *customerView) add() {
	fmt.Println("--------------------添加客户-------------")
	v.customerService.Add()
	fmt.Println("--------------------添加客户完成-------------")
}

//删除客户
func (v *customerView) delete() {
	fmt.Println("------------删除客户------------")
	v.customerService.Delete()
}

//修改客户
func (v *customerView) update() {
	fmt.Println("------------修改客户------------")

	v.customerService.Update()
	fmt.Println("------------修改客户完成------------")
}

//退出
func (v *customerView) exit() {
	fmt.Print("你确定要退出吗？(y/n): ")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("你的输入有误，请重新输入(y/n)： ")
	}
	if choice == "y" {
		v.loop = false
	}

}

//显示主菜单
func (v customerView) mainMenu() {
	for {
		fmt.Println("------------客户信息管理软件------------")
		fmt.Println("             1 添 加 客 户")
		fmt.Println("             2 修 改 客 户")
		fmt.Println("             3 删 除 客 户")
		fmt.Println("             4 客 户 列 表")
		fmt.Println("             5 退	出")
		fmt.Print("请选择(1-5): ")
		fmt.Scanln(&v.key)
		switch v.key {
		case "1":
			v.add()
		case "2":
			v.update()
		case "3":
			v.delete()
		case "4":
			v.list()
		case "5":
			v.exit()
		default:
			fmt.Println("你的输入有误。请重新输入: ")
		}

		if !v.loop {
			break
		}
		fmt.Println()
	}
	fmt.Println("你已经退出客户关系管理系统")

}

func main() {
	customerView := customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(), //一开始就初始化一个客户
	}

	customerView.mainMenu()
}
