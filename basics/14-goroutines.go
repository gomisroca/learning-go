package main

import (
	"fmt"
	"time"
)

func exFunc(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {
	exFunc("direct")

	// A goroutine is a lightweight thread managed by the Go runtime.
	// It is a function that runs concurrently with other goroutines.
	// We can start a goroutine by using the go keyword followed by a function call.
	go exFunc("goroutine")
	
	// We can also start a goroutine with an anonymous function.
	go func(msg string) {
		fmt.Println(msg)
	}("going")


	// The non-goroutine function will execute synchronously,
	// meaning it will block the main function until it finishes.
	// The goroutines will run concurrently, meaning it will not block the main function.
	// Meaning, both goroutines will execute concurrently.

	// The main function will continue to run concurrently with the goroutines.
	// If the main function exits before the goroutines finish, they will be terminated.
	// To prevent the main function from exiting immediately, we can use time.Sleep to wait.

	time.Sleep(time.Second)
	fmt.Println("done")
}