package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [5]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		random := rand.Intn(100) //[0,100)
		arr[i] = random
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}

	fmt.Println(arr)
}
