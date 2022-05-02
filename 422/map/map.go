package main

import "fmt"

func main() {

	map1 := make(map[int]string)
	map1[0] = "13"
	map1[5] = "56456"
	map1[6] = "123"
	map1[7] = "865"

	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Printf("key：%v, value: %v\n", k, v)
	}
	//所以map内部还是无序存储的，只是fmt.println时，自动按照key的顺序输出

}
