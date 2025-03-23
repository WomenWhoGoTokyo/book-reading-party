// P.152
package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.InitADC()
	sensor := machine.ADC{Pin: machine.WIO_LIGHT}
	sensor.Configure(machine.ADCConfig{})
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		val := sensor.Get()
		fmt.Printf("%04X\n", val)
		// 暗いとき、AD値が0x8000より小さい(1.65V以下)ときにLEDを点灯する
		if val < 0x3000 {
			led.High()
		} else {
			led.Low()
		}
		time.Sleep(50 * time.Millisecond)
	}
}
