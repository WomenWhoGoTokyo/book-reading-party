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
	// 初期化
	i2c.WriteRegister(0x18, 0x20, []byte{0x57})
	data := []byte{0, 0, 0, 0, 0, 0}
	for {
		// XYZ軸の重力加速度を読み出し
		i2c.ReadRegister(0x18, 0x28|0x80, data)
		x := readAcceleration(data[0], data[1])
		y := readAcceleration(data[2], data[3])
		z := readAcceleration(data[4], data[5])
		fmt.Printf("X:%6.2f Y:%6.2f Z:%6.2f\r\n", x, y, z)
		time.Sleep(100 * time.Millisecond)
	}
}
func readAcceleration(l, h byte) float32 {
	// uint8の値を組み合わせてuint16型の変数を作成
	a := uint16(l) | uint16(h)<<8
	// 0x4000 == 1gのため、0x4000で割る
	return float32(int16(a)) / 0x4000
}
