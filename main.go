package main

import (
	"machine"
	"time"
)

const SLEEP_TIME = time.Millisecond * 1000

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Low()
		time.Sleep(SLEEP_TIME)

		led.High()
		time.Sleep(SLEEP_TIME)
	}
}
