package main

import (
	"fmt"
	"time"
)

// Select lets us wait on multiple channels at the same time.


func channelSelect() {
    // Init channels
    c1 := make(chan string)
    c2 := make(chan string)

    // Simulate some time passing in an operation, as it might in a real world scenario
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // We await both of these values simultaneously with select
    // When one of the channels receives a value, it will be printed
    // Essentially, it "selects" the first channel that receives a value, or one at random if multiple channels are ready
    // Sort of like a "switch" statement for channels

    // Since we are using a for loop, it will run select until the for loop is done
    // Without the for loop, we would only get the first value to arrive to either of the channels, in this case "one"
    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
        // We can also use a default case, which will be executed if no other case is executed
        // This is ufel to try a send or receive without blocking
        // default:
        //     fmt.Println("no message received")
    }
}