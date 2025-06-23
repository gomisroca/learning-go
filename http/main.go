package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	pwd := r.FormValue("password")
	
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	if pwd == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	// Here we would typically check the credentials against a database or other storage
	// For this example, we will just check against hardcoded values
	// In a real application, you would have a salted hash for the password and check it securely
	if id != "admin" || pwd != "secret" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// If the credentials are valid, we can proceed with the login
	fmt.Fprintf(w, "Login successful!\n")
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// This handler would typically handle token refresh logic
	// For simplicity, we will just demand a token and return a success message
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	// Here you would typically validate the token and refresh it
	if token != "valid-token" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// If the token is valid, we can proceed with the refresh
	// In a real application, you would generate a new token and return it
	// For this example, we will just return a new random token
	newToken := "new-valid-token"
	fmt.Fprintf(w, "Token refreshed successfully! New token: %s\n", newToken)
}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/refresh-token", refreshTokenHandler)
	http.ListenAndServe(":8080", nil)
}