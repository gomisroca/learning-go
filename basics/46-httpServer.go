package main

import (
	"fmt"
	"net/http"
)

// Handlers implement the http.Handler interface
// they take a ResponseWriter and a Request
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// Handlers can read everything from the request, including headers
func headersHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func httpServer() {
	// Register our handlers on routes in the default router
	http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/headers", headersHandler)

	// Serve the default router on port 8080
	http.ListenAndServe(":8080", nil)
}