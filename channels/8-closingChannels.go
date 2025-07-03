package main

import "fmt"

func closingChannels() {
	// Closing a channel indicates that no more values will be sent into it.

	jobs := make(chan int, 5)
	done := make(chan bool)

	// Goroutine to receive from the jobs channel
	go func() {
		for {
			// more will be false if the channel is closed and no more values will be received
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send some jobs
	for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
	// Close the jobs channel
    close(jobs)
    fmt.Println("sent all jobs")

	// Wait for the goroutine to finish
	<-done

	// Reading from a closed channel will return the zero value for the channel's type
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}