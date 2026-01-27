// RWMutex = Read-Write Mutex, it's a mutex with two kinds of locks:
// - Write lock - exclusive (same as Mutex) and
// - Read lock - shared (many readers allowed)
// Use it when there are lots of reads, but few writes
// Rules:
// 1. Multipe RLock() are allowed
// 2. Lock() blocks while any readers exists
// 3. RLock() blocks while a write exists

package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]int),
	}
}

// Read (shared lock)
func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.m[key]
	return val, ok
}

// Write (exclusive lock)
func (c *Cache) Set(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = val
}

func RWMutex() {
	cache := NewCache()
	ctx, cancel := context.WithCancel(context.Background())

	// Write goroutine
	go func() {
		for i := range 5 {
			cache.Set("count", i)
			fmt.Println("Set: ", i)
			time.Sleep(500 * time.Millisecond)
		}
		cancel() // tell readers to stop
	}()

	// Multiple reader goroutines
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Reader %d exits\n", i)
					return
				default:
					if val, ok := cache.Get("count"); ok {
						fmt.Printf("Reader %d got: %d\n", id, val)
					}
				}
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All readers exited")
}
