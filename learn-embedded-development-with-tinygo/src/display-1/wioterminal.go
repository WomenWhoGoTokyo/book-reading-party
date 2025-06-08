//go:build wioterminal
// +build wioterminal

package main

import (
	"machine"

	"github.com/sago35/tinydisplay/examples/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

var (
	button machine.Pin
)

func initEnv() *ili9341.Device {
	button = machine.WIO_5S_PRESS
	button.Configure(machine.PinConfig{Mode: machine.PinInput})
	return initdisplay.InitDisplay()
}
func Pressed() bool {
	return !button.Get()
}
