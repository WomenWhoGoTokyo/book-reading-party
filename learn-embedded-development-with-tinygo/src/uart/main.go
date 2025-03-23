// P.147
package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	ser := machine.UART1
	ser.Configure(machine.UARTConfig{
		BaudRate: 9600,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	})
	time.Sleep(2 * time.Second) // USB CDC接続待ち
	// UARTに送信する
	fmt.Fprintf(ser, "message-to-uart\r\n")
	// UARTから受信する
	msg := ""
	fmt.Fscanf(ser, "%s\r\n", &msg)
	// USBCDCに送信する
	fmt.Printf("%s\r\n", msg)
}
