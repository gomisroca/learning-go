package main

import (
	"fmt"
	"time"
)

func tickers() {
	// Similar to timers, but instead of doing something once in the future,
	// We do something repeatedly at a given interval.

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Goroutine that will print the tick at the given interval
	// if the ticker is stopped, it will stop
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// We stop the ticker after 1.6 seconds, so it should tick 3 times
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}