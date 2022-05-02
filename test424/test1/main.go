package main

import "fmt"

//1)编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个10*8的矩形，在main方法中调用该方法。
type MethodUtils struct {
}

func (mu MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	mu := MethodUtils{}
	mu.print()
}
