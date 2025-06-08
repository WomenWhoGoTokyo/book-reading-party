package main

import (
	"machine"
	"machine/usb/adc/midi"
	"time"

	"tinygo.org/x/drivers/tone"
)

func main() {
	speaker, _ := tone.New(machine.TCC0, machine.BUZZER_CTR)
	notes := []tone.Note{tone.A5, tone.B5, tone.C6, tone.D6, tone.E6, tone.F6, tone.G6}

	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m := midi.New()
	m.SetHandler(func(b []byte) {
		// 外部からのMIDI信号を受信
		led.Toggle()
		switch b[1] {
		case 0x90:
			// on
			note := notes[0]
			switch midi.Note(b[2]) {
			case midi.C4:
				note = notes[2]
			case midi.D4:
				note = notes[3]
			case midi.E4:
				note = notes[4]
			default:
			}
			speaker.SetNote(note)
		case 0x80:
			// off
			speaker.Stop()
		}
	})
	time.Sleep(1 * time.Second)

	for {
		// 外部へMIDI信号を送信
		m.NoteOn(0, 1, midi.C4, 0x7F)
		time.Sleep(time.Millisecond * 1000)
		m.NoteOff(0, 1, midi.C4, 0x7F)
		time.Sleep(time.Millisecond * 1000)
	}
}
