package main

import (
	"fmt"
	"time"
)

func timeouts() {
    c1 := make(chan string, 1)
	// We create a mock external call that will return its result after 2 seconds
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
	// We add a timeout of 1 second, 
	// so if the external call takes longer than that, we will get a timeout
	// In this case, since the external call took 2 seconds, we will get the timeout
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
	// Here, the timeout is 3 seconds, so we will get the external call result
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}