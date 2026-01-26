// select lets a goroutine wait on multiple channel operations and proceed with one that is ready.

package goroutine

import (
	"context"
	"fmt"
	"sync"
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

func FanInChannels() { // Multiplexing channels using select and for loop
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Producer goroutines
	go func() {
		for i := range 5 {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := range 3 {
			ch2 <- fmt.Sprintf("Msg %d", i)
			time.Sleep(150 * time.Millisecond)
		}
		close(ch2)
	}()

	// Consumer loop using select to read from both channels
	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil // Set to nil to disable this case
			} else {
				fmt.Println("Received from ch1:", val)
			}
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println("Received from ch2:", msg)
			}
		}
	}

	fmt.Println("All channels closed, exiting loop.")
}

func WorkerGracefulShutdownWithDoneSignal() {
	jobs := make(chan int)
	done := make(chan bool)

	// Worker goroutine
	go func() {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					fmt.Println("Jobs channel closed.")
				} else {
					fmt.Println("Processing job:", job)
					time.Sleep(500 * time.Millisecond) // Simulate work
				}
			case <-done:
				fmt.Println("Received done signal, worker exiting.")
				return // No leaked goroutine
			}
		}
	}()

	// Send jobs
	for i := range 5 {
		jobs <- i
	}
	close(jobs)  // Close jobs channel to signal no more jobs
	done <- true // Send done signal
	time.Sleep(1 * time.Second)
}

func WorkerGracefulShutdownWithContext() {
	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	// Worker goroutine
	go func() {
		defer wg.Done()
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					fmt.Println("Jobs channel closed.")
				} else {
					fmt.Println("Processing job:", job)
					time.Sleep(500 * time.Millisecond) // Simulate work
				}
			case <-ctx.Done():
				fmt.Println("Received done signal, worker exiting.")
				return // No leaked goroutine
			}
		}
	}()

	// Send jobs
	for i := range 5 {
		jobs <- i
	}
	close(jobs) // Close jobs channel to signal no more jobs
	cancel()    // Cancel context to signal done

	wg.Wait()
}
