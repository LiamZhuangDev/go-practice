// A fixed number of worker goroutines that pull jobs from a shared queue and process them concurrently.

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Job int

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
	}
}

func WorkerPoolTest() {
	const workerCount = 3
	const jobCount = 10
	jobs := make(chan Job)
	var wg sync.WaitGroup

	// start workers
	for i := range workerCount {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// submit jobs
	for j := range jobCount {
		jobs <- Job(j)
	}
	close(jobs)

	wg.Wait()
	fmt.Println("all workers done")
}
