package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		fmt.Println()
		for j := 1; j <= i; j++ {
			sum := i * j
			fmt.Printf("%v * %v = %v \t", j, i, sum)
		}
	}
}
