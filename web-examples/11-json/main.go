package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// We will see how to encode and decode JSON data using the encoding/json package.

// Struct we will use to encode and decode JSON data.
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// Decode JSON data.
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User // Struct where we will decode JSON data to.
		json.NewDecoder(r.Body).Decode(&user) // Decode JSON data into user.

		// NOTE
		// Since we added `json:"firstname"` etc. to the struct,
		// we will access the fields using those names, not the struct tags.
		// i.e user.firstname, not user.Firstname

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})

	// Encode JSON data.
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		// Struct we want to encode to JSON data.
		john := User{
            Firstname: "John",
            Lastname:  "Doe",
            Age:       25,
		}
		json.NewEncoder(w).Encode(john) // Encode john to JSON data.

		// NOTE
		// Similar to above, our fields will be turned to use our json tags.
		// i.e. firstname, not Firstname
	})

	http.ListenAndServe(":8080", nil)
}