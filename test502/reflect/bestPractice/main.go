package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	//获取reflect.Type、reflect.Value和a的对应类别
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//判断是不是结构体，如果不是结构体，就退出
	if kd != reflect.Struct {
		fmt.Printf("expect struct\n")
		return
	}
	//获取到这个结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d 值为=%v\n", i, val.Field(i)) //val.Filed(i) 获取到第几个字段，从0开始
		//获取到struct标签！ 注意需要使用通过reflect.Type来获取tag标签的值！
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有标签就显示，否则不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为: %v\n", i, tagVal)
		}
	}
	//获取该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	//调用的是第二个方法，0是第一个。  方法按照方法名的字母排序
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

/*
	使用反射遍历结构体的字段，调用结构体的方法，获取结构体标签的值
*/
func main() {
	var a = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}
