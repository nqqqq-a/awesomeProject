package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	//功能一：客户端发送单行数据，然后退出  os.Stdin获得标准输入【终端】
	reader := bufio.NewReader(os.Stdin)

	for {
		//代表读取一行
		rs, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerString err=", err)
			return
		}
		rs = strings.Trim(rs, " \r\n")
		if rs == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		//再将rs发送给服务器,参数得是[]byte切片
		_, err = conn.Write([]byte(rs + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
			return
		}
	}
}
