package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	stringChan := make(chan string, 5)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("我是%v", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据为%v\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据为%v\n", v)
		default:
			fmt.Printf("都读取不到了，不玩了")
			return
		}
	}

}
