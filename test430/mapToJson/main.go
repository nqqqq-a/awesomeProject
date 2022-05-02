package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("结构体序列化结果：%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("map 序列化结果：%v\n", string(data))
}

func testSlice() {
	//切片里有多个map
	var slice []map[string]interface{}
	//slice = make([]map[string]interface{}, 4)
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = []string{"北京", "上海"}

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = []string{"深圳", "广州"}
	slice = append(slice, m2)

	//对切片序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("slice 序列化结果：%v\n", string(data))

}

func testFloat64() {
	var num1 float64 = 2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("对普通类型 序列化结果：%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
