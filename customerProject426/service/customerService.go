package service

import (
	"awesomeProject/customerProject426/model"
	"fmt"
)

//该customer完成对customer的操作(增删改查)
type CustomerService struct {

	//customers切片
	customers []model.Customer
	//当前有多少个用户,也可以作为新客户的字段
	customerNum int
}

//获取customerService实例
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerservice := &CustomerService{}
	customerservice.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "13823376950", "945096934@qq.com")
	customerservice.customers = append(customerservice.customers, customer)
	return customerservice
}

//返回客户切片
func (vs *CustomerService) List() []model.Customer {
	return vs.customers
}

func (vs *CustomerService) Add() bool {
	fmt.Print("姓名： ")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别： ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄： ")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话号码： ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱： ")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)
	vs.customerNum++
	//确定一个分配id的规则，就是添加的顺序
	customer.SetId(vs.customerNum)

	vs.customers = append(vs.customers, customer)

	return true
}

//根据id返回一个customer对应的下标
func (vs *CustomerService) findById(id int) int {

	for i := 0; i < len(vs.customers); i++ {
		if vs.customers[i].GetId() == id {
			return i
		}
	}
	return -1
}

//根据id从切片中删除
func (vs *CustomerService) Delete() {
	fmt.Print("请选择要删除的客户编号(-1退出): ")
	var id int
	fmt.Scanln(&id)
	//输入-1时退出删除
	if id == -1 {
		return
	}

	//获得下标
	index := vs.findById(id)
	if index == -1 {
		fmt.Println("不存在这个id客户...")
		return
	}

	fmt.Print("确认是否删除(Y/N)：")
	isDelete := ""
	fmt.Scanln(&isDelete)

	//当输入N或n时，不进行删除了
	if isDelete == "N" || isDelete == "n" {
		return
	}

	vs.customers = append(vs.customers[:index], vs.customers[index+1:]...)
}

func (vs *CustomerService) Update() {
	fmt.Println("请选择待修改的客户编号(-1退出)： ")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	//获取下标
	i := vs.findById(id)

	fmt.Printf("姓名(%v)： ", vs.customers[i].Name)
	name := ""
	fmt.Scanln(&name)
	vs.customers[i].Name = name

	fmt.Printf("性别(%v)： ", vs.customers[i].Gender)
	gender := ""
	fmt.Scanln(&gender)
	vs.customers[i].Gender = gender

	fmt.Printf("年龄(%v)： ", vs.customers[i].Age)
	age := 0
	fmt.Scanln(&age)
	vs.customers[i].Age = age

	fmt.Printf("电话号码(%v)： ", vs.customers[i].Phone)
	phone := ""
	fmt.Scanln(&phone)
	vs.customers[i].Phone = phone

	fmt.Printf("邮箱(%v)： ", vs.customers[i].Email)
	email := ""
	fmt.Scanln(&email)
	vs.customers[i].Email = email

}
