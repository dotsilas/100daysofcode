package main

import (
)

func main() {
	// channel creation: make( chan keyword and data type to transport )
	channel := make(chan string)
	go func() {
		channel <- "Hello world"
	}()

	message := <- channel
	println(message)
}

