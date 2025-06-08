package main

import (
	"machine"
	"machine/usb/adc/midi"
	"time"
)

func main() {
	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	m := midi.New()

	previousState := true
	for {
		currentState := button.Get()

		if previousState && !currentState {
			velocity := uint8(0x50)

			m.NoteOn(0, 1, midi.G4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.G4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.A4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.A4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.B4, velocity)
			time.Sleep(400 * time.Millisecond)
			m.NoteOff(0, 1, midi.B4, velocity)
			time.Sleep(200 * time.Millisecond)

			m.NoteOn(0, 1, midi.A4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.A4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.G4, velocity)
			time.Sleep(300 * time.Millisecond)
			m.NoteOff(0, 1, midi.G4, velocity)
			time.Sleep(200 * time.Millisecond)

			m.NoteOn(0, 1, midi.G4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.G4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.A4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.A4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.B4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.B4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.A4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.A4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.G4, velocity)
			time.Sleep(200 * time.Millisecond)
			m.NoteOff(0, 1, midi.G4, velocity)
			time.Sleep(50 * time.Millisecond)

			m.NoteOn(0, 1, midi.A4, velocity)
			time.Sleep(400 * time.Millisecond)
			m.NoteOff(0, 1, midi.A4, velocity)

			time.Sleep(500 * time.Millisecond)
		}

		previousState = currentState
		time.Sleep(10 * time.Millisecond)
	}
}
