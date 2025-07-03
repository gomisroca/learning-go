package main

import "fmt"

// When we use a channel as a func argument, we can specify the direction of the channel.

// For example, here we can ONLY send to the channel, not receive from it.
// If we try to receive from the channel inside this function, we will get a compile error.
func ping(pings chan<- string, msg string) {
	pings <- msg // Send to pings channel
}

// And here, we can ONLY receive from the pings channel,
// and we can ONLY send to the pongs channel.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // Receive from pings channel
	pongs <- msg   // Send to pongs channel
}

func channelDirections() {
	pings := make(chan string, 1) // Init a channel with a buffer of 1
	pongs := make(chan string, 1)

	ping(pings, "passed message") // We will send "passed message" to the pings channel
	pong(pings, pongs) // We will receive "passed message" from the pings channel and send it to the pongs channel

	fmt.Println(<-pongs) // We block until we receive a message from the pongs channel
}