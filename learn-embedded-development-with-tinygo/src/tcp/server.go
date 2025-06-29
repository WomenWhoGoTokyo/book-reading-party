package main

import (
	"fmt"
	"log"
	"net"
)

var (
	serverIP = "192.168.11.2"
	port     = 8088
)

func main() {
	conn, err := net.Listen(`tcp`, fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("listen: %s:%d\n", serverIP, port)
	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Printf("from client: %q\r\n", string(buf[:n]))
	conn.Write(buf[:n])
	conn.Close()
}
