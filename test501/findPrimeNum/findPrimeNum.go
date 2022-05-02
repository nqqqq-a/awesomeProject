package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func findPrimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		time.Sleep(1000)
		num, ok := <-intChan

		if !ok {
			break
		}
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}

	}
	fmt.Println("取不到数据了退出")
	exitChan <- true
}

func putNum(intChan chan int) {
	for i := 1; i <= 100; i++ {
		intChan <- i
	}
	close(intChan)
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	runtime.GOMAXPROCS(cpu)
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	//需求:要求统计1-200000的数字中，哪些是素数?

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go findPrimeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		close(primeChan)
	}()

	i := 1
	for res := range primeChan {
		fmt.Printf("第%v个素数：%v\n", i, res)
		i++
	}
}
