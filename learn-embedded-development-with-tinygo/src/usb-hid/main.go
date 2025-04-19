package main

import (
	"machine"
	"machine/usb/hid/mouse"
	"math"
	"time"
)

func main() {
	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	m := mouse.New()

	for {
		if !button.Get() {
			const radius = 150
			const steps = 100
			const sleep = 2 * time.Millisecond
			const offsetAngle = -math.Pi / 2

			points := [5][2]int{}
			for i := 0; i < 5; i++ {
				angle := offsetAngle + 2*math.Pi*float64(i)/5
				x := int(math.Cos(angle) * radius)
				y := int(math.Sin(angle) * radius)
				points[i] = [2]int{x, y}
			}

			seq := []int{0, 2, 4, 1, 3, 0}

			m.Press(mouse.Left)
			for i := 0; i < len(seq)-1; i++ {
				from := points[seq[i]]
				to := points[seq[i+1]]
				px := from[0]
				py := from[1]
				for j := 1; j <= steps; j++ {
					x := from[0] + (to[0]-from[0])*j/steps
					y := from[1] + (to[1]-from[1])*j/steps
					m.Move(x-px, y-py)
					px = x
					py = y
					time.Sleep(sleep)
				}
			}
			m.Release(mouse.Left)

			m.Move(20, 20)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
