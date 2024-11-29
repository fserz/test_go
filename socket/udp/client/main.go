package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Printf("dial failed, err is %v\n", err)
		return
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	s, err := input.ReadString('\n')
	if err != nil {
		fmt.Printf("read string failed, err is: %v\n", err)
		return
	}
	n, err := conn.Write([]byte(s))
	if err != nil {
		fmt.Printf("send to server failed, err: %v\n", err)
		return
	}
	fmt.Printf("send data length is %d\n", n)
	// 接收数据
	var buf [1024]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Printf("receive from UDP failed, err: %v\n", err)
		return
	}
	fmt.Printf("read from %v, message: %v\n", addr, string(buf[:n]))

}
