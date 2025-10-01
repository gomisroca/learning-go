package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Sometimes we need our program to handle Unix signals
// i.e. a server shutting down gracefully after SIGTERM
func signals() {
	// Go signal notifications work by sending os.Signal values to a channel
	sigs := make(chan os.Signal, 1)

	// signal.Notify registers a channel to receive the specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive from sigs in the main fn, but we'll do it in a goroutine

	done := make(chan bool, 1)

	// executes a blocking receive on sigs
	// will print out the received signal and notify the program to exit
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}