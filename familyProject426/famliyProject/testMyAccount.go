package main

import "fmt"

func main() {

	//声明一个变量，接受用户的选择
	key := ""
	//loop 控制for循环
	loop := true

	//定义账户余额 (默认10000元)
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//每次收支的详情
	details := "收支\t账户金额\t收支金额\t说   明"

	//循环显示主菜单
	for {

		fmt.Println("--------------家庭收支记账软件-------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Print("               请选择(1-4)：")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("--------------当前收支明细-------------")
			fmt.Println(details)
		case "2":
			fmt.Println("--------------登记收入-------------")
			fmt.Print("本次收入金额: ")
			fmt.Scanln(&money)
			balance += money
			fmt.Print("本次收入说明: ")
			fmt.Scanln(&note)
			//将这次的收入情况拼接到details变量中
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)

		case "3":
			fmt.Println("--------------登记支出-------------")
			fmt.Print("本次支出金额: ")
			fmt.Scanln(&money)

			if money > balance { //判断余额是否足够
				fmt.Println("余额不足")
				break
			} else {
				balance -= money
			}

			fmt.Print("本次支出说明: ")
			fmt.Scanln(&note)
			//将这次的支出情况拼接到details变量中
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)

		case "4": //退出软件
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
				loop = false
			}

		default:
			fmt.Println("请输入正确的选项..")
		}

		fmt.Println()

		if !loop {
			break
		}

	}
	fmt.Println("你退出了家庭收支记账软件...")
}
