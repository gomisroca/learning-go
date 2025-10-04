package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

// html/template provides templating for HTML.
func main() {
	// We parse the HTML template from the file
	tmpl := template.Must(template.ParseFiles("layout.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// The data passed to the template can be any Go struct
		// In this case, we're passing a TodoPageData struct we created
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
		// Execute the template in the response writer
        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":80", nil)
}