// A fixed number of worker goroutines that pull jobs from a shared queue and process them concurrently.

package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job int

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("job canceled, worker %d exiting\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("jobs closed, worker %d exiting\n", id)
				return
			}
			fmt.Printf("worker %d processing job %d\n", id, job)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func WorkerPoolTest() {
	const workerCount = 3
	const jobCount = 10
	jobs := make(chan Job)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start workers
	for i := range workerCount {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// cancel jobs after 1 second
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	// submit jobs
jobLoop:
	for j := range jobCount {
		select {
		case jobs <- Job(j):
			fmt.Printf("submitted job %d\n", j)
		case <-ctx.Done():
			fmt.Println("canceling jobs...")
			break jobLoop
		}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("all workers done")
}
