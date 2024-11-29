package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
	}
	defer conn.Close()
	for {
		var buf [128]byte
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read fromm udp failed, err: %v\n", err)
		}
		fmt.Println("receive data: ", string(buf[:n]))
		lenByte, err := conn.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed, err is: %v\n", addr, err)
			return
		}
		fmt.Println("write to udp data length is: ", lenByte)
	}
}
