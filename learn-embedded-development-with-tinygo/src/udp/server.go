package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", 8088))
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listen: :%d\n", 8088)
	buf := [1024]byte{}
	for {
		n, _, _ := conn.ReadFromUDP(buf[:])
		fmt.Printf("from client: %q\r\n", string(buf[:n]))
	}
}
