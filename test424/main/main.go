package main

import (
	"awesomeProject/test424/model"
	"fmt"
)

func main() {
	var student = model.NewStudent("nq~", 88.8)
	fmt.Println(student)
	fmt.Println(*student)
	fmt.Println(student.GetName())
}
