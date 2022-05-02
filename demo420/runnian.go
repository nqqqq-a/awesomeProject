package main

import "fmt"

func main() {
	var year int
	fmt.Scanln(&year)
	//if year%400 == 0 {
	//	fmt.Println("yes")
	//} else {
	//	if year%4 == 0 {
	//		if year%100 != 0 {
	//			fmt.Println("yes")
	//		} else {
	//			fmt.Println("no")
	//		}
	//	}
	//}

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
