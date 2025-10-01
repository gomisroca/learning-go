package main

import (
	"fmt"
	"net/http"
)

// A basic HTTP server should be able to
// Handle dynamic requests
// Serve static assets
// Accept connections

// We can handle requests with net/http
// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
//     fmt.Fprint(w, "Welcome to my website!")
// })
// http.Request contains information about the request, such as GET parameters
// r.URL.Query().Get("token")
// or POST parameters (from a form)
// r.FormValue("email")

// We can serve static assets with http.FileServer
// fs := http.FileServer(http.Dir("static/"))

func main() {
	// Handle requests to the root path
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	// Serve static assets from the static/ directory
    fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
	
}