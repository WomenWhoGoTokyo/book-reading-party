package main

import (
	"image/color"
	"math/rand"
	"time"
)

func main() {
	display := initEnv()
	rand.Seed(time.Now().UnixNano())
	for {
		if Pressed() {
			display.FillScreen(color.RGBA{uint8(rand.Uint32()),
				uint8(rand.Uint32()), uint8(rand.Uint32()), 255})
			time.Sleep(100 * time.Millisecond)
		}
	}
}
