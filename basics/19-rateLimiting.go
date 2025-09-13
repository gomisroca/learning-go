package main

import (
	"fmt"
	"time"
)

// Important mechanism to control resource utilization.
// Go does it with goroutines, channels and tickers.
func rateLimiting() {
	// Basic rate limiting
	// Let's say we want to limit our handling of incoming requests

	// Serve the requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// We use a 1 second ticker that will act as the limiter
	limiter := time.Tick(time.Second)

	for req := range requests {
		// We block on a receive from the limiter
		// Limiting ourselves to 1 req every second
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Let's say we want to allow a burst of requests
	// We can accomplish this by buffering the limiter channel
	burstyLimiter := make(chan time.Time, 3)

	// Fill the limiter channel with 3 values, so up to 3 requests can happen right away
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Add a new value to the limiter channel every second, if the channel is not full
	// Eventually, if no requests are received, the limiter chan will get full again
	// That would allow another burst of requests
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	// Simulate 6 requests
	burstyRequests := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		// We block on a receive from the limiter
		// Limiting ourselves to 3 inmediate requests, then 1 req every second
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// Output:
	// Basic rate limit @ 1s:
	// request 1 2025-09-13 21:08:40.0823221 +0200 CEST m=+1.000639401
	// request 2 2025-09-13 21:08:41.0818515 +0200 CEST m=+2.000168801
	// request 3 2025-09-13 21:08:42.0823602 +0200 CEST m=+3.000677501
	// request 4 2025-09-13 21:08:43.0822897 +0200 CEST m=+4.000607001
	// request 5 2025-09-13 21:08:44.0822019 +0200 CEST m=+5.000519201
	// Bursty rate limit (notice burst of 3):
	// request 1 2025-09-13 21:08:44.0822019 +0200 CEST m=+5.000519201
	// request 2 2025-09-13 21:08:44.0822019 +0200 CEST m=+5.000519201
	// request 3 2025-09-13 21:08:44.0822019 +0200 CEST m=+5.000519201
	// request 4 2025-09-13 21:08:45.0827599 +0200 CEST m=+6.001077201
	// request 5 2025-09-13 21:08:46.0825706 +0200 CEST m=+7.000887901
	// request 6 2025-09-13 21:08:47.0825454 +0200 CEST m=+8.000862701
}