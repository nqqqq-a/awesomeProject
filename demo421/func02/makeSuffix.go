package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name没有指定后缀，则加上，否则返回原来名字
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
}
func main() {
	f := makeSuffix(".jpg")
	fmt.Println("文件处理后=", f("winter"))
}
