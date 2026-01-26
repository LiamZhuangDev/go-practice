package goroutine

import (
	"fmt"
	"sync"
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

func FanOutWorkersWithBufferedChannel() {
	ch := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(2)

	// Consumers
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Routine1 received: ", v)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Routine2 received: ", v)
			time.Sleep(time.Second)
		}
	}()

	// Producer
	for i := range 40 {
		ch <- i
		fmt.Println("Sent: ", i)
	}

	close(ch) // Producer closes the channel to signal no more values will be sent
	wg.Wait() // Wait for both consumers to finish
}
