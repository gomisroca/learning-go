package main

import "fmt"

func nonblocking() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	// If a value is available on messages, it will take the <-messages case,
	// otherwise, it will take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send
	// Here msg cannot be sent into messages because the channel is unbuffered and has no receiver
	// Therefore, the select will take the default case
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}
	
	// Multi-way non-blocking select
	// We can have multiple cases in the select
	// Here, we attempt non-blocking receives on both channels
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}