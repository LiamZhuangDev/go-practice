package concurrency

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int64
}

func NewCounter() *Counter {
	return &Counter{
		mu:  sync.Mutex{},
		val: 0,
	}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.val++
}

func (c *Counter) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.val
}

func SafeCounterTest() {
	c := NewCounter()
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Printf("Count: %d\n", c.Get())
}
