package goroutine

import (
	"fmt"
	"time"
)

// An unbuffered channel provides synchronous communication between goroutines.
// A send on an unbuffered channel blocks until another goroutine is ready to receive from that channel, and vice versa.
func UnbufferedChannelExample() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		ch <- "Hello"
		ch <- "World"
	}()

	time.Sleep(1000 * time.Millisecond)

	for msg := range ch {
		fmt.Println("Received: ", msg)
	}
}

// A buffered channel provides asynchronous communication between goroutines.
// A send on a buffered channel blocks only when the buffer is full, and a receive blocks only when the buffer is empty.
func BufferedChannelExample() {
	ch := make(chan string, 2) // Buffered channel with capacity of 2, which allows 2 messages to be sent without blocking.

	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "From Buffered Channel"
		close(ch)
	}()

	time.Sleep(1000 * time.Millisecond)

	for msg := range ch {
		fmt.Println("Received: ", msg)
	}
}
