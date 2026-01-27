// If we start goroutines and there is a request to cancel or a timeout happens, we need a standard and safe way to tell all related goroutines to stop
// That's what context.Context is for.
// context.Context is an object that carries:
// 1. Cancellation signal
// 2. Deadline / timeout
// 3. Request-scoped values
// And it propagates across (flows through) goroutines.
package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	// Worker Goroutine
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Received done signal, exiting.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Cancel the processing
	time.Sleep(time.Second)
	fmt.Println("Cancelling.")
	cancel()

	wg.Wait()
}

func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timed out, exiting")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func ContextWithDeadline() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Deadline passes, exiting...")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
}
