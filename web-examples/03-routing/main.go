package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// One thing net/http doesn't do very well is complex routing.
// We will use a library called gorilla/mux to handle this.

func main() {
	// Create a new request router
	// It will be the main router for our application
	r := mux.NewRouter()

	// Now we can register request handlers like usual
	// Only difference is we do r.HandleFunc instead of http.HandleFunc
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	// Big positive of mux is the ability to extract parameters from the URL
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // get the route segments
		title := vars["title"] // the book title slug
        page := vars["page"] // the page

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// We can restrict the methods that can be used on a route
	r.HandleFunc("/get-only", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This route is only accessible via GET")
	}).Methods("GET")

	// We can also restrict the routes to specific HTTP/HTTPS
	// And we can restrict the routes to specific hosts or subdomains
	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com").Schemes("https")

	// We can add subrouters with a prefix
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "Book: %s\n", title)
	}) // /books/{title}

	// Finally, add the router as a "main router" param of http.ListenAndServe
	http.ListenAndServe(":80", r)
}