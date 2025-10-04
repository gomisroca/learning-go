package main

import (
	"fmt"
	"log"
	"net/http"
)

// Basic logging middleware

// The middleware takes a http.HandlerFunc as a parameter
// Wraps it and returns a new http.HandlerFunc
func loggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// Some example route handlers
func foo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "bar")
}

func main() {
	// We will use the loggingMiddleware to wrap our route handlers
	http.HandleFunc("/foo", loggingMiddleware(foo))
	http.HandleFunc("/bar", loggingMiddleware(bar))

	http.ListenAndServe(":8080", nil)
}