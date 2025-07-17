package main

import "fmt"

func rangeOverChannels() {
	// Using range, we can iterate over the values in a channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range iterates over each element in the channel as it is received
	// since the channel is closed, range will stop after the second element is received
	for elem := range queue {
		fmt.Println(elem)
	}
}