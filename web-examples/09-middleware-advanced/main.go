package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Again the middleware takes a http.HandlerFunc as a parameter
// Wraps it and returns a new http.HandlerFunc

// We will define a new type Middleware
// It will make it easier to chain multiple middlewares together

type Middleware func(http.HandlerFunc) http.HandlerFunc


// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			// Call the next middleware/handler
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, 
// else returns a 400 Bad Request
func Method(m string) Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
            }

			// Call the next middleware/handler
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Example route handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	// We will use the Chain function to apply multiple middlewares to our route handler
	// If we do a non-GET request, we will get a 400 Bad Request
	// We will log the request path and the time it took to process
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}