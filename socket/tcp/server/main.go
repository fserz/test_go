package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// 针对当前连接做数据的发送
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n ", err)
			return
		}
		recv := string(buf[:n])
		fmt.Println("receive data: ", recv)
		conn.Write([]byte("ok seal ball"))
	}
}

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err : %+v", err)
		return
	}
	for {
		// 2. 等待客户端连接
		conn, err := listener.Accept()
		fmt.Println("listener address is ", listener.Addr())
		if err != nil {
			fmt.Printf("accepted failed, err: %v\n", err)
		}
		// 3. 启动一个goroutine处理连接
		go process(conn)
	}
}
