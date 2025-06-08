//go:build !baremetal
// +build !baremetal

package main

import (
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var display *initdisplay.TinyDisplay

func initEnv() *initdisplay.TinyDisplay {
	display = initdisplay.InitDisplay()
	return display
}
func Pressed() bool {
	return display.GetPressedKey() != 0xFFFF
}
