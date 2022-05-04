package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis 写入数据和读取数据

	//1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close() //一定要关闭
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//2.通过 go向redis写入数据 string[k-v]
	_, err = conn.Do("Set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.通过go读取redis数据string[k-v]
	//因为返回的result是interface{} 类型，所以要转换redis.String
	result, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	fmt.Printf("%v", result)
	fmt.Println()

}
