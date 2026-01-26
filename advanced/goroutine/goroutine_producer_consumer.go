package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int) {
	for i := range 20 {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, id int, wg *sync.WaitGroup) { // wg is shared state, must be shared among consumers by passing WaitGroup pointer
	defer wg.Done()

	for val := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, val)
		time.Sleep(50 * time.Millisecond)
	}
}

func ProducerConsumer() {
	ch := make(chan int, 5)
	numConsumers := 3

	var wg sync.WaitGroup
	wg.Add(numConsumers)

	go producer(ch)

	for i := range numConsumers {
		go consumer(ch, i, &wg)
	}

	wg.Wait() // Why not wg->Wait()? Methods are defined on pointer receivers internally
}
