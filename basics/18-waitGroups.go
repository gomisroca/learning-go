package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a WaitGroup.

// Function that will run in every goroutine.
func wrkr(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter by one.
		go func(id int) {
			wrkr(id)
			wg.Done() // Decrement the WaitGroup counter by one.
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish.
}