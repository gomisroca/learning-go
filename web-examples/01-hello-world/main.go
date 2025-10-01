package main

import (
	"fmt"
	"net/http"
)

// Go has a built-in web server with the net/http package

// First, we create a handler function
// func (w http.ResponseWriter, r *http.Request)

func main() {
	// With http.HandleFunc, we register a handler function to a route
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// To listen for HTTP connections, we use http.ListenAndServe
	http.ListenAndServe(":80", nil)
	// Now if we go to http://localhost/ we should see "Hello, you've requested: /"
}