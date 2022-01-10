package main

import (
	"blinky/worker"
	"fmt"
	"time"
)

func main() {
	status := make(chan bool)
	go worker.Work(status)

	for {
		select {
		case <-status:
			break
		default:
			fmt.Println("Waiting...\r")
			time.Sleep(time.Millisecond * 1)
		}
	}
}
