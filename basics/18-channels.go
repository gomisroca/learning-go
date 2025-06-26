package main

import "fmt"

// Channels are used to communicate between goroutines.

// Can send values into channels in one goroutine,
// and receive those values in another goroutine.

func channels() {
	// To create a channel, we use the make() function with the chan keyword.
    messages := make(chan string)

	// The channel <- operator is used to send a value into a channel.
    // Here, we send "ping" into the messages channel.
	go func() { messages <- "ping" }()

	// The <- channel operator is used to receive a value from a channel.
	// Here, we receive the "ping" value from the messages channel.
    msg := <- messages

	// By default, sends and receives block until both the sender and receiver are ready. 
	// Meaning we can wait for the "ping" here with no problem.
    fmt.Println(msg)
}