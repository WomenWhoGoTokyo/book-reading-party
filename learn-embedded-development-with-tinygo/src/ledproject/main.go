package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		println("Hello World", "\r")
		led.Toggle()
		time.Sleep(1000 * time.Millisecond)
	}
}
