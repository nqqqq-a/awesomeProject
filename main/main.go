package main

import (
	"awesomeProject/Sort"
	"fmt"
)

func main() {
	fmt.Println("hello goland")
	arr := []int{4, 6, 84, 46, 56, 45, 5, 15}
	Sort.BubbleSort(arr)
	fmt.Println(arr)

	arr2 := []int{9, 12, 84, 46, 56, 45, 5, 15, 55}
	Sort.QuickSort(arr2, 0, len(arr2)-1)
	fmt.Println(arr2)
}
