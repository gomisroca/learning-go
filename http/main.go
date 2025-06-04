package http

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Your Token is: %s\n", r.URL.Query().Get("token"))
}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}