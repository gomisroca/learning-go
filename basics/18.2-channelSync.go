package main

import (
	"fmt"
	"time"
)

// We can use a channel to synchronize goroutines.
// Here, we create a channel (done) with a buffer size of 1.
// Then, we start a goroutine (worker) that sends a value into the done channel.
// Finally, we receive the value from the done channel in the main goroutine.

// The main goroutine will block until the worker goroutine sends a value into the done channel.

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSync() {
	done := make(chan bool, 1)

	go worker(done)

	// Without this, the main goroutine would exit before worker even started.
	<-done
}