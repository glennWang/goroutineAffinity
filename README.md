## goroutineAffinity

inspired by [CPU Affinity in Go](http://pythonwise.blogspot.com/2019/03/cpu-affinity-in-go.html)

### usage
set goroutine affinity to specified cpu core

### support
* Linux

### todo
* support MacOS
* support Windows

### How to use
```
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/glennWang/goroutineAffinity"
)

func randSleep() {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
}

func worker(id int, lock bool) {
	if lock {
		goroutineAffinity.SetAffinity(id)
	}

	for {
		fmt.Printf("worker: %d, CPU: %d\n", id, goroutineAffinity.GetAffinityCPU())
		randSleep()
	}
}

func main() {

	fmt.Println("NumCPU:", runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, true)
	}
	time.Sleep(2 * time.Second)
}

```

### ⚠️ License

Source code in `goroutineAffinity` is available under the [MIT License](/LICENSE).

### 🤝 Welcome pr