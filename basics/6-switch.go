package main

import (
	"fmt"
	"runtime"
	"time"
)

func switches() {
	fmt.Print("Go runs on ")

	// Just like if statements or for loops, we can make a small statement before the switch block
	// in this case, "os := runtime.GOOS", which will only be available in the switch block
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("macOS.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}

	// We can also omit the condition
	// This serves as a cleaner way to write if-else statements
	t:= time.Now()
	switch {
		case t.Hour() < 12:  // if t.Hour() < 12
			fmt.Println("Good morning.")
		case t.Hour() < 16: // else if t.Hour() < 16
			fmt.Println("Good afternoon.")
		default: // else
			fmt.Println("Good evening.")
	}
}