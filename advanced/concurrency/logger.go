package concurrency

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func NewLogger(filename string) *Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	l := &Logger{
		ch: make(chan string, 100),
	}

	// write logs
	go func() {
		defer file.Close()

		for msg := range l.ch {
			func() {
				defer l.wg.Done() // avoid panics that could cause WaitGroup to block forever
				defer func() {    // recover panic so logger goroutine keep running
					if r := recover(); r != nil {
						fmt.Println("logger panic recovered, ", r)
					}
				}()

				if _, err := fmt.Fprintln(file, msg); err != nil {
					fmt.Println("write log failed, ", err)
				}
			}()
		}
	}()

	return l
}

func (l *Logger) Log(msg string) {
	l.wg.Add(1)
	l.ch <- msg
}

func (l *Logger) Close() {
	l.wg.Wait()
	close(l.ch)
}

func LoggerTest() {
	fmt.Printf("[%v] start logger test\n", now())

	logger := NewLogger("app.log")

	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range 3 {
				logger.Log(fmt.Sprintf("[goroutine %d] log %d", id, j))
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()

	logger.Close()

	fmt.Printf("[%v] all logs written\n", now())
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
