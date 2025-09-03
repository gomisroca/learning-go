package main

import (
	"fmt"
	"time"
)

// Go has built-in support for timers.
// time.NewTimer builds a channel that will send a value after the specified duration.
func timers() {
	timer1 := time.NewTimer(2 * time.Second)

	// We can use the timer channel to wait for the timer to expire.
	<-timer1.C
    fmt.Println("Timer 1 fired")

	// To wait, we can simply use time.Sleep,
	// however, timer are useful because they can be stopped.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timers time to fire, if they are going to
	time.Sleep(2 * time.Second)

	// The first timer will fire, but the second will not since it was stopped before it had a chance to fire.
}