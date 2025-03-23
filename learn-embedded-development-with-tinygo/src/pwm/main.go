package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/tone"
)

func main() {
	speaker, _ := tone.New(machine.TCC0, machine.BUZZER_CTR)
	notes := []tone.Note{tone.A5, tone.B5, tone.C6, tone.D6, tone.E6, tone.F6, tone.G6}
	i := 0
	for {
		speaker.SetNote(notes[i])
		time.Sleep(100 * time.Millisecond)
		i = (i + 1) % len(notes)
	}
}
