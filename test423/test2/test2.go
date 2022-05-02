package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	cat1 := Cat{"小白", 12, "white"}
	var cat2 Cat
	cat2.Name = "小黑"
	cat2.Age = 10
	cat2.Color = "black"

	fmt.Println(cat1, cat2)

}
