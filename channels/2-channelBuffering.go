package main

import "fmt"

// By default, channels are unbuffered.
// This means they only accept sends ( chan <- ) if there is a corresponding receive ( <- chan ) ready.

// Buffered channels can accept a limited number of values without a receiver.
func channelBuffering() {
	messages := make(chan string, 2) // buffer size of 2

	// Since messages is buffered, we can send two values into it without a corresponding receive.
	messages <- "buffered"
	messages <- "channel"

	// Then later on, we can receive the two values.
	fmt.Println(<- messages)
	fmt.Println(<- messages)
}