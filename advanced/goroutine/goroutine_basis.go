// A goroutine is a lightweight concurrent execution unit managed by Go’s runtime (not an OS thread).

package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// When demo/main exits, the program exits immediately,
// The goroutine may not get CPU time before main ends
// time.Sleep is a hack to keep main alive long enough for the goroutine to run.
// This is not production-safe — it’s just for demos.
func demo() {
	go sayHello()
	time.Sleep(time.Second)
}

func WaitGroupExampe() {
	// “How to wait for all goroutines to finish?”
	// wg.Add(1) increments the counter by 1
	// wg.Done() decrements the counter by 1
	// wg.Wait() blocks until the counter is 0
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d\n", id)
		}(i) // pass i as argument to avoid closure capture issue
	}

	wg.Wait()
}
