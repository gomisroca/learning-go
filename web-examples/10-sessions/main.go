package main

// How to store data in session cookies
// We will use the gorilla/sessions package
import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Create a new cookie-based session store
// Session data is stored in encrypted cookies
// Very similar to JWT
var (
	// Key must be 16, 24 or 32 bytes long
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name") // Get the session

    // Check if user is authenticated
	// This will "decrypt" the cookie and check the "authenticated" value
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    // Print secret message
    fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name") // Get the session

	// Here we would authenticate the user
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name") // Get the session

	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
    http.HandleFunc("/secret", secret)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":8080", nil)
}