// A goroutine is a lightweight concurrent execution unit managed by Go’s runtime (not an OS thread).

package goroutine

import (
	"fmt"
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
