package main

import "fmt"

//2)编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形

type MethodUtils struct {
	m, n int
}

func (mu MethodUtils) print(m, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	mu := MethodUtils{5, 10}
	mu.print(mu.m, mu.n)
}
