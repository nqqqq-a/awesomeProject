package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	n := 0
	personChan := make(chan *Person, 10)
	for i := 0; i < 10; i++ {
		person := Person{
			Name:    fmt.Sprintf("我叫做%v", n),
			Age:     n,
			Address: fmt.Sprintf("我住在%v号", n),
		}
		personChan <- &person
		n++
	}
	close(personChan)

	//for person := range personChan {
	//	fmt.Println(person)
	//}

	for person := range personChan {
		fmt.Println(person.Name, "\t我的年龄是", person.Age, "\t", person.Address)
	}

}
