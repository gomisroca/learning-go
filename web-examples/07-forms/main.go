package main

import (
	"net/http"
	"text/template"
)

// We will simulate a contact form
// And parse the message into a struct

// Message struct
type ContactDetails struct {
	Email    string
	Subject  string
	Message  string
}

func main() {
	// We will use a template to render the form
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We will check if the request method is POST
		if r.Method != http.MethodPost {
			// If not, we will render the form
			tmpl.Execute(w, nil)
			return
		}

		// We will parse the form values into a ContactDetails struct
		details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // Do something with details
        _ = details

		// If everything is ok, we will render a success message
        tmpl.Execute(w, struct{ Success bool }{true})
	})
	
    http.ListenAndServe(":8080", nil)
}