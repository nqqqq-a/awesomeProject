package utils

import "fmt"

type familyAccount struct {
	//声明一个变量，接受用户的选择
	key string
	//loop 控制for循环
	loop bool

	flag bool
	//定义账户余额 (默认10000元)
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//每次收支的详情
	details string
}

//将显示明细写成一个方法
func (this *familyAccount) showDetails() {
	fmt.Println("--------------当前收支明细-------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细。。。来一单吧")
	}

}

//登记收入的方法
func (this *familyAccount) income() {
	fmt.Println("--------------登记收入-------------")
	fmt.Print("本次收入金额: ")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("本次收入说明: ")
	fmt.Scanln(&this.note)
	//将这次的收入情况拼接到details变量中
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//登记支出的方法
func (this *familyAccount) pay() {
	fmt.Println("--------------登记支出-------------")
	fmt.Print("本次支出金额: ")
	fmt.Scanln(&this.money)

	if this.money > this.balance { //判断余额是否足够
		fmt.Println("余额不足")
		return
	} else {
		this.balance -= this.money
	}

	fmt.Print("本次支出说明: ")
	fmt.Scanln(&this.note)
	//将这次的支出情况拼接到details变量中
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//退出的方法
func (this *familyAccount) exit() {
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
		this.loop = false
	}

}

//编写一个工厂模式的构造方法,返回*FamilyAccount实例
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		flag:    false,
		balance: 10000,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说   明",
	}
}

//给这个结构体编写方法
//显示主菜单
func (this *familyAccount) MainMenu() {
	for {
		fmt.Println("--------------家庭收支记账软件-------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Print("               请选择(1-4)：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4": //退出软件
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
		fmt.Println()
		if !this.loop {
			break
		}
	}
}
