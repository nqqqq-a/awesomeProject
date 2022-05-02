package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"牛魔拳\"}"
	var master Monster

	err := json.Unmarshal([]byte(str), &master)
	if err != nil {
		fmt.Printf("反序列化失败，err是 %v", err)
	}
	fmt.Printf("反序列化后 master"+
		" = %v\n", master)

}
func main() {
	unmarshalStruct()
}
