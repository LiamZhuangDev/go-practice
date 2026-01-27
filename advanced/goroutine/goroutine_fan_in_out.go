package goroutine

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producerWithContext(ctx context.Context, wg *sync.WaitGroup, id int, ch chan<- int) {
	defer wg.Done()

	for {
		val := rand.Intn(1000)

		select {
		case ch <- val:
			fmt.Printf("producer %d -> %d\n", id, val)
		case <-ctx.Done():
			fmt.Printf("producer %d exiting\n", id)
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func consumerWithContext(ctx context.Context, wg *sync.WaitGroup, id int, ch <-chan int) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Printf("consumer %d <- %d\n", id, val)
		case <-ctx.Done():
			fmt.Printf("consumer %d exiting\n", id)
			return
		}
	}
}

func FanInOut() {
	ch := make(chan int, 5)
	ctx, cancel := context.WithCancel(context.Background())

	numProducers := 2
	numConsumers := 3

	var wg sync.WaitGroup
	wg.Add(numProducers + numConsumers)

	for i := range numProducers {
		go producerWithContext(ctx, &wg, i, ch)
	}

	for i := range numConsumers {
		go consumerWithContext(ctx, &wg, i, ch)
	}

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("Cancel")

	wg.Wait()
}
