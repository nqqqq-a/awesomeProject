package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个Hero结构体的切片
type HeroSlice []Hero

//3.实现Interface接口
func (h HeroSlice) Len() int {
	return len(h)
}

//Less方法决定使用什么标准排序
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

func (h HeroSlice) Swap(i, j int) {
	//temp := h[i]
	//h[i] = h[j]
	//h[j] = temp
	h[i], h[j] = h[j], h[i] //Go里面的交换写法可以这样
}

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对 intSlice排序（使用系统提供的方法）
	//sort.Ints(传入一个切片)  可以进行一个排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	//使用sort.Sort(data Interface)  ： 该interface要实现Len(), Less(), Swap()
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			fmt.Sprintf("英雄~%d", rand.Intn(100)),
			rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println(heros)

	sort.Sort(heros)

	fmt.Println(heros)
}
