package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.11.60", 8088))
	if err != nil {
		log.Fatal(err)
	}
	// サーバーに送信する
	fmt.Fprintf(conn, "hello from go\r\n")
	// サーバーからの受信を標準出力に書き出す
	io.Copy(os.Stdout, conn)
	conn.Close()
}
