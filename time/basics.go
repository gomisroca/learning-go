package main

import (
	"fmt"
	"time"
)

func timeBasics() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

	// We can extract the usual components of a time struct
	p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
	p(then.Weekday())

	// We can compare times
	p(then.Before(now))
	p(then.Equal(now))
	p(then.After(now))

	// .Sub returns the duration between two times
	diff := now.Sub(then)
	p(diff)

	// We can get the duration in various units
	p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

	// We can add and subtract durations
	p(then.Add(diff))
	p(then.Add(-diff))
}