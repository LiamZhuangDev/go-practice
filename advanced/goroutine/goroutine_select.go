// select lets a goroutine wait on multiple channel operations and proceed with one that is ready.

package goroutine

import (
	"fmt"
	"time"
)

func SelectBasis() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	time.Sleep(1500 * time.Millisecond)

	// Use select to wait on multiple channel operations
	// The first channel that is ready will proceed
	// If both are ready, one is chosen at random
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}

func SelectNonBlockingWithDefault() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
			return
		default:
			fmt.Println("No message received yet, doing other work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func SelectWithTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Hello from channel"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received: ", msg)
	case <-time.After(2 * time.Second): // time.After returns a receive-only channel (<-Chan Time) that will send the current time (as timeout signal) after the specified duration
		fmt.Println("Timeout: No message received within 2 seconds")
	}
}
