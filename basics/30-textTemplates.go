package main

import (
	"os"
	"text/template"
)

// Go has built-in support to create dynamic content or showing custom output

func textTemplates() {
	// We create a new template and parse its body from a string
	t1 := template.New("t1")
	// this {{.}} is an action, used to dynamically insert content
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	// Alternatively use Must, which will panic if there is an error during Parse
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text") // the {{.}} is replaced with "some text"
	t1.Execute(os.Stdout, 5) // the {{.}} is replaced with 5
	t1.Execute(os.Stdout, []string{
		"Go",
        "Rust",
        "C++",
        "C#",
	}) // the {{.}} is replaced with the array

	// Helper fn we will use below
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	// If the data is a struct, we can use dot notation to access its fields
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"}) // the {{.Name}} is replaced with "Jane Doe"

	// We can also use the dot notation to access map values
	 t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })

	// if else
	// a value is considered false is its the default value for its type
	// - trims whitespace
	t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
	// ie. if "not empty" is a non-empty string, print yes, else no
	t3.Execute(os.Stdout, "not empty") // yes
	t3.Execute(os.Stdout, "") // no

	// in range blocks, {{.}} is the current value of the iteration
	t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}