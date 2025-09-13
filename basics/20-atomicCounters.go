package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Although the main way of managing state in Go is through channels,
// there are other ways to do it.
func atomicCounters() {
	// Integer representing our counter
	var ops atomic.Uint64

	// WaitGroup helper to wait for all goroutines to finish
	var wg sync.WaitGroup

	// 50 goroutines that will increment the counter
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Atomically increment the counter by 1, 1000 times
				ops.Add(1)
			}
		})
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// We will get exactly 50,000
	fmt.Println("ops:", ops.Load())
}