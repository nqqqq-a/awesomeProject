package main

import (
	"fmt"
	"reflect"
)

func test1(b interface{}) {
	//获取类型
	rTyp := reflect.TypeOf(b)

	fmt.Println("rTyp:", rTyp)

	rValue := reflect.ValueOf(b)
	n2 := 2 + rValue.Int()
	fmt.Println(n2)
	fmt.Println("rValue:", rValue)

	iv := rValue.Interface()
	num2 := iv.(int)
	fmt.Println(num2)
}

type Student struct {
	name string
	age  int
}

func test2(b interface{}) {
	//获取类型
	rTyp := reflect.TypeOf(b)

	fmt.Println("rTyp:", rTyp)

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue.kind:", rValue.Kind())
	fmt.Println("rValue:", rValue)

	iv := rValue.Interface()
	fmt.Printf("iv= %v, type = %T \n", iv, iv)
	stu := iv.(Student)
	fmt.Println(stu.name)
}

func main() {
	//编写一个案例，演示对(基本数据类型！、interface{}、reflect.Value)反射的操作
	var num int = 100
	test1(num)

	fmt.Println("-------")
	//演示对(结构体！、interface{}、reflect.Value)反射的操作
	var stu = Student{
		name: "nq",
		age:  22,
	}

	test2(stu)
}
