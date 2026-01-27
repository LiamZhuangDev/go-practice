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

// context.WithValue is for request-scoped metadata that flows through APIs, not for passing real data.
// This example follows handle - service - repo structure
type reqIDKeyType struct{} // unexported, unique key

var reqIDKey = reqIDKeyType{}

func handler(ctx context.Context) {
	// attach value
	ctx = context.WithValue(ctx, reqIDKey, "req-123")

	service(ctx)
}

func service(ctx context.Context) {
	repo(ctx)
}

func repo(ctx context.Context) {
	// retrieve value
	if id, ok := ctx.Value(reqIDKey).(string); ok {
		fmt.Println("request id: ", id)
	}
}

func ContextWithValue() {
	handler(context.Background())
}

// context.WithValue, together with cancellation
func handler2(ctx context.Context) {
	// attach value
	ctx = context.WithValue(ctx, reqIDKey, "req-456")

	// add cancellation
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	fmt.Println("handler start")

	err := service2(ctx)
	if err != nil {
		fmt.Println("handler exit: ", err)
	}

	fmt.Println("handler done")
}

func service2(ctx context.Context) error {
	if id, ok := ctx.Value(reqIDKey).(string); ok {
		fmt.Println("service request id: ", id)
	}

	return repo2(ctx)
}

func repo2(ctx context.Context) error {
	if id, ok := ctx.Value(reqIDKey).(string); ok {
		fmt.Println("repo request id: ", id)
	}

	for range 5 {
		select {
		case <-ctx.Done():
			fmt.Println("repo canceled: ", ctx.Err())
			return ctx.Err()
		default:
			fmt.Println("repo working...")
			time.Sleep(400 * time.Millisecond)
		}
	}

	return nil
}

func ContextWithValueAndCancel() {
	handler2(context.Background())
}

// CascadeCancellation
func ContextWithCascadeCancel() {
	parentCtx, parentCancel := context.WithCancel(context.Background())

	child1Ctx, child1Cancel := context.WithCancel(parentCtx)
	defer child1Cancel()

	child2Ctx, child2Cancel := context.WithCancel(parentCtx)
	defer child2Cancel()

	go func() {
		time.Sleep(1 * time.Second)
		parentCancel()
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker1 canceled")
				return
			default:
				fmt.Println("Worker1 working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(child1Ctx)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker2 canceled")
				return
			default:
				fmt.Println("Worker2 working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(child2Ctx)

	wg.Wait()
}
