## Concurrency

### Share by Communicating

Go encourages sharing values on channels, and never actively shared by separate threads of execution. Only one goroutine has access to the value at any given time. Data races cannot occur.

### Goroutines

A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. If one is blocked, the others can continue.

A goroutine is created by calling `go` followed by a function call:

```go
go list.Sort() // run list.Sort concurrently; don't wait for it
```

In Go, function literals are closures, the variables they reference are bound to the function's environment.

### Channels

Channels are allocated with the `make` function. The resulting value acts as a reference to an underlying data structure.

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

With a channel, we can give commands to a goroutine:

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

A buffered channel is like a semaphore, they might limit throughput.

```go
var sem = make(chan int, MaxOutstanding) // Buffered channel with "MaxOutstanding" max amount of values.

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

A better approach would be to start a fixed number of `handle` goroutines reading from the request channel. This will limit the number of simultaneous requests.

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}
```

### Channels of Channels

A channel is a first-class value and can be allocated and passed around like any other value. For example earlier we didn't define how Request looks:

```go
type Request struct {
    args []int
    f func([]int) int
    resultChan chan int
}
```

It provides a function and its arguments, and a channel to receive the result into:

```go
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
```
