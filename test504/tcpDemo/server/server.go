package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里循环接收客户端发送的数据
	defer conn.Close()
	for {
		//创建一个新切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息，如果客户端没有write【发送】。则协程在此阻塞
		//fmt.Printf("服务器等待客户端%s 发送信息...\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端Read err=", err)
			return
		}
		//显示客户端发送的内容到服务端
		fmt.Printf(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	//1. tcp表示使用网络协议是tcp
	//2. 0.0.0.0：8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close() //延时关闭listen

	//循环等待客户端来连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			fmt.Printf("Accept() suc con = %v, 服务器ip=%v 客户端ip=%v\n", conn, conn.LocalAddr(), conn.RemoteAddr())
		}

		//这里准备其中一个协程，为客户端服务（这样可以多个客户端一起发送数据）
		go process(conn)
	}

	//fmt.Printf("listen suc=%v\n", listen)

}
