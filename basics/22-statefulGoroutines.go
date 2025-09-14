package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// We can use the built-in sync features of goroutines and channels to manage state

// Our state will be owned by a single goroutine
// Other goroutines will send operations to the stateful goroutine
// These types encapsulate the operations
type readOp struct {
	key int // The key to read
	resp chan int // Channel receiving the response
}

type writeOp struct {
	key int // The key to write
	val int // The value to write
	resp chan bool // Channel receiving the response
}

func statefulGoroutines() {
	var readOps uint64
	var writeOps uint64

    // Other goroutines will issue read/write requests through these channels
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Stateful goroutine
	go func() {
		var state = make(map[int]int) // The state, a map of keys to values
		for { // Loop forever, selecting on the channels
			select {
			case read := <-reads: // Wait for a readOp to come into the reads chan
				read.resp <- state[read.key] // Send response to the readOp's resp channel
			case write := <-writes: // Wait for a writeOp to come into the writes chan
                state[write.key] = write.val // Update the state
				write.resp <- true // Send response to the writeOp's resp channel
			}
		}
	}()

    // 100 goroutines that will issue read requests
	for range 100 {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 10 goroutines that will issue write requests
	for range 10 {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // Let goroutines run for a second
    time.Sleep(time.Second)

    // Capture and report the operation counts
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}