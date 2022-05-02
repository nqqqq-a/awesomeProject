package main

import (
	"fmt"
	"runtime"
)

//1)启动一个协程，将1-2000的数放入到一个channel中 ,比如numChan
//2)启动8个协程，从numChan取出数(比如n)，并计算1+..+n的值，并存放到resChan
//3)最后8个协程协同完成工作后，再遍历resChan,显示结果 [如res[1]= 1 .. res[10]= 55..]
//4)注意:考虑resChan chan int是否合适?

func cal(numChan chan int, resChan chan map[int]int, exitChan chan bool) {

	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		rm := make(map[int]int)
		for i := 1; i <= n; i++ {
			res += i
		}
		rm[n] = res
		resChan <- rm
	}
	exitChan <- true
}

func main() {

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores - 1)

	numChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 2000)
	exitChan := make(chan bool, 8)

	//启动一个协程，将1-2000的数放入到一个channel中 ,比如numChan
	go func() {
		for i := 1; i < 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	//2)启动8个协程，从numChan取出数(比如n)，并计算1+..+n的值，并存放到resChan
	for i := 0; i < 8; i++ {
		go cal(numChan, resChan, exitChan)
	}

	//这里没读出数据的时候，会堵塞，所以得8次的写入都完成，这个exitChan的8次读取才会运行结束
	for i := 0; i < 8; i++ {
		<-exitChan
	}
	close(resChan)

	for v := range resChan {
		for i, value := range v {
			fmt.Printf("res[%v] = %v\n", i, value)
		}
	}

}
