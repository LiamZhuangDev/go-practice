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

	time.Sleep(3000 * time.Millisecond)

	for msg := range ch {
		fmt.Println("Received: ", msg)
	}
}
