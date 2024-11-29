package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1. 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	defer conn.Close()
	if err != nil {
		fmt.Printf("dial failed, err: %v\n", err)
		return
	}
	// 2. 利用连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		str, err := input.ReadString('\n')
		if err != nil {
			fmt.Printf("read string err, err is %v\n", err)
		}
		str = strings.TrimSpace(str)
		if strings.ToUpper(str) == "Q" {
			return
		}
		// 给客户端发消息
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("send failed, err: %v\n", err)
			return
		}
		fmt.Printf("write %d byte to client\n", n)
		// 从服务端接收回复的消息
		var buf [1024]byte
		n, err = conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed, err %v\n", err)
			return
		}
		fmt.Println("receive from server: ", string(buf[:n]))
	}
}
