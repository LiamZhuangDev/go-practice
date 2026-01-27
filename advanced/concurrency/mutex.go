// Mutex = Mutual Exclusion, it's used to protect shared data so that only one goroutine can access it at a time
// So a mutex prevents data race

package concurrency

import (
	"fmt"
	"sync"
)

// var counter int

// func increment() {
// 	counter++ // Data race if multiple goroutines increment the counter at the same time
// }

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()

	counter++
}

func MutexExample() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
