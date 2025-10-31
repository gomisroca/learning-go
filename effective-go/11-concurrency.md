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
