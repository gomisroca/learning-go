package main

import (
	"fmt"
	"time"
)

func formattingAndParsing() {
	p := fmt.Println

	t := time.Now()
	// we can use .Format() to format a time.Time
	p(t.Format(time.RFC3339))

	// we can parse a time.Time from a string, giving it the format we want
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// We can use examples to have it format the time for us
	// The examples must always follow Mon Jan 2 15:04:05 MST 2006
	// So if we want the current time as 9:20PM, we need 3:04PM as the example
	p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)

	// We can also use standard string formatting
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

	// Parse returns an error on malformed input
	ansic := "Mon Jan _2 15:04:05 2006"
    _, e := time.Parse(ansic, "8:41PM")
    p(e)
}