package main

import "fmt"

func main() {
	str := "helloNq"
	arr1 := []rune(str)
	arr1[5] = '聂'
	fmt.Println(str)
	fmt.Println(string(arr1))
	str = string(arr1)
	fmt.Println(str)
}
