package task13

import (
	"fmt"
	"sync"
)

func Start() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i) // не передали ссылку на WaitGroup из main рутины
	}
	wg.Wait()
	fmt.Println("exit")
}
