package main

import (
	"fmt"
	"sync"
)

// Although the main way of managing state in Go is through channels,
// there are other ways to do it.
// Atomic Counters was one of them, and mutexes are another.

// Mutex (mutual exclusion) only allow one goroutine to access a variable at a time
// In this case, we will use it to protect the counters map.
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

// Lock the mutex, increment the counter and unlock the mutex
func (c *Container) inc(name string) {
	c.mu.Lock() // Lock so only one goroutine can access the map at a time
	defer c.mu.Unlock() // Unlock when we are done via defer
	c.counters[name]++
}

func mutexes() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Increment a counter n times
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}
	
	// Run several goroutines concurrently
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println(c.counters)
}
