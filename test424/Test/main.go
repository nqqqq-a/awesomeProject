package main

import "fmt"

//定义小小计算器结构体(calculator) ,实现加减乘除四个功能I
//实现形式:用一个方法搞定，需要接收两个数，还有一个运算符

type Calculator struct {
	Num1, Num2 float64
}

func (ca *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = ca.Num1 + ca.Num2
	case '-':
		res = ca.Num1 - ca.Num2
	case '/':
		res = ca.Num1 / ca.Num2
	case '*':
		res = ca.Num1 * ca.Num2
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}

func main() {
	ca := Calculator{4.0, 6.2}
	fmt.Println(ca.getRes('-'))
}
