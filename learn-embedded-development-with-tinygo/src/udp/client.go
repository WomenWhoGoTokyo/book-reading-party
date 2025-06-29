package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("udp", fmt.Sprintf("255.255.255.255:%d", 8088))
	for i := 0; i < 3; i++ {
		fmt.Printf("send to 255.255.255.255:%d\n", 8088)
		fmt.Fprintf(conn, "udp from go %d\r\n", i)
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}
