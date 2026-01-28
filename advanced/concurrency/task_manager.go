// Task Manager supports multi-tasking, cancellation and timeout.
// The APIs contains:
// NewTaskManager() - constructor
// Start(string, task) - start a task
// StartWithTimeout(string, task, time.Duration) - start a task with timeout
// CancelAll() - cancel all running tasks
// Wait() - wait all tasks complete

package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task func(ctx context.Context) error

type TaskManager struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewTaskManager(timeout time.Duration) *TaskManager {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return &TaskManager{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (tm *TaskManager) Start(name string, task Task) {
	tm.wg.Add(1)

	go func() {
		defer tm.wg.Done()

		if err := task(tm.ctx); err != nil {
			fmt.Printf("[%s] failed: %v\n", name, err)
			return
		}

		fmt.Printf("[%s] completed successfully\n", name)
	}()
}

func (tm *TaskManager) StartWithTimeout(name string, task Task, timeout time.Duration) {
	tm.wg.Add(1)

	go func() {
		defer tm.wg.Done()

		ctx, cancel := context.WithTimeout(tm.ctx, timeout)
		defer cancel()

		if err := task(ctx); err != nil {
			fmt.Printf("[%s] failed: %v\n", name, err)
			return
		}

		fmt.Printf("[%s] compeleted successfully\n", name)
	}()
}

func (tm *TaskManager) CancelAll() {
	tm.cancel()
}

func (tm *TaskManager) Wait() {
	tm.wg.Wait()
}

func computeTask(ctx context.Context) error {
	sum := 0
	for i := range 1_000_000 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			sum += i
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("compute sum: %d\n", sum)
		}
	}
	fmt.Printf("compute result: %d\n", sum)
	return nil
}

func ioTask(ctx context.Context) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := range 10 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			fmt.Printf("io task processsing chunk %d\n", i)
		}
	}
	return nil
}

func TaskManagerTest() {
	tm := NewTaskManager(3 * time.Second)

	tm.Start("compute", computeTask)
	tm.StartWithTimeout("io", ioTask, time.Second)

	cancel := true
	if cancel {
		time.Sleep(2 * time.Second)
		tm.CancelAll()
	} else {
		// timeout after 3 seconds or earlier (timeout specified in StartWithTimeout)
	}

	tm.Wait()
	fmt.Println("All tasks done")
}
