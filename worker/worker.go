package worker

import (
	"machine"
	"time"
)

const SLEEP_TIME = time.Millisecond * 1

var led machine.Pin

func init() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func Blink() {
	led.Low()
	time.Sleep(SLEEP_TIME)

	led.High()
	time.Sleep(SLEEP_TIME)
}

func Work(status chan bool) {
	for i := 0; i <= 48000; i++ {
		Blink()
	}

	close(status)
}
