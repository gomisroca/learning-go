package main

import (
	"fmt"
	"time"
)

// Getting the time since Unix epoch
func epochs() {
	now := time.Now()
    fmt.Println(now)

	fmt.Println(now.Unix()) // seconds since epoch
	fmt.Println(now.UnixMilli()) // milliseconds since epoch
	fmt.Println(now.UnixNano()) // nanoseconds since epoch

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
