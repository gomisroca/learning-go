package main

import (
	"fmt"
	"time"
)

// We will create a worker, and run several concurrent instances of it.
// The worker will receive tasks in the jobs channel and send the results to the results channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second) // Simulate an expensive task
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerPools() {
	const numJobs = 5

	// Init the channels that will send tasks and receive results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send numJobs amount of tasks to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel to indicate that no more tasks will be sent

	// Collect the results from the results channel
	for a := 1; a <= numJobs; a++ {
		<-results
		// If we want to print out the results:
   		// result := <-results  
		// fmt.Println("result:", result)
	}

	// The output should be something like this:
	// worker 3 started job 2
	// worker 2 started job 3
	// worker 1 started job 1
	// worker 1 finished job 1
	// worker 1 started job 4
	// worker 2 finished job 3
	// worker 2 started job 5
	// worker 3 finished job 2
	// worker 2 finished job 5
	// worker 1 finished job 4

	// In essence, the 3 workers are running concurrently, and as they finish their tasks, they send the results and then take on a new task from the jobs channel.
}