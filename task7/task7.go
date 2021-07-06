package task7

import (
	"fmt"
	"sync"
	"time"
)

func writeMap(myMap map[int]int, i int, m *sync.RWMutex) {
	m.Lock()
	myMap[i]++
	m.Unlock()
}

func Start() {
	myMap := map[int]int{}
	m := sync.RWMutex{}

	for i := 0; i < 5; i++ {
		go writeMap(myMap, i, &m)
	}

	time.Sleep(1 * time.Millisecond)

	m.RLock()
	fmt.Println("results: ", myMap)
	m.RUnlock()
}
