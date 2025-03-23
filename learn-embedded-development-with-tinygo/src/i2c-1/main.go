package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL0_PIN,
		SDA: machine.SDA0_PIN,
	})
	time.Sleep(2 * time.Second) // USB CDC接続待ち
	// 読み出し用のスライスを用意して読み込む
	data := []byte{0}
	err := i2c.ReadRegister(0x18, 0x0F, data)
	if err != nil {
		// エラー処理を行う
	}

	for {
		fmt.Printf("Who am I : 0x%02X\r\n", data[0])
		time.Sleep(time.Second)
	}
	select {}
}
