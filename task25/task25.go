package task25

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Println(c.value)
}

func Start() {
	counter := Counter{value: 1}

	fmt.Println(counter.value)
	go counter.Increment()
	go counter.Increment()
	go counter.Increment()
	go counter.Increment()
	go counter.Increment()

	time.Sleep(time.Second)
}
