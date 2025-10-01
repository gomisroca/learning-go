package main

import (
	"fmt"
	"net/http"
	"time"
)

// HTTP Servers are useful for demonstrating context.Context for controlling cancellation
// Context carries deadlines, cancelation signals, etc. across API boundaries.

func helloHdlr(w http.ResponseWriter, req *http.Request) {
	// Every request creates a new context
	ctx := req.Context() // access the context of the request
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler finished")

	select {
	// Send something back to the client after 10s
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	// Check if the context has been canceled/timed out/etc
	case <-ctx.Done():
		err := ctx.Err() // Will give us info about why the context was canceled
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func context() {
	http.HandleFunc("/hello", helloHdlr)
	http.ListenAndServe(":8080", nil)
}