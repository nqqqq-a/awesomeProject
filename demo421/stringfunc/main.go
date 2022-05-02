package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello"
	println(str, len(str))

	str = "hello聂"
	println(str, len(str)) //可以发现长度变成8了，因为go的string按照utf-8，汉字占3个字节

	//要用[]rune切片 解决汉字存在的string类型的长度计算
	// str2 := []rune(str)
	// str2就会按照字符去遍历，借用for{}
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}
	println(len(str2))

	//str3 := "456"
	r, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(r)
	}

}
