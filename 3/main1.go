package main

import (
	"fmt"
	"sync"
)

func sumSquaresNumbers(num int, sum *int, wg *sync.WaitGroup, m *sync.Mutex) {
	square := num * num

	m.Lock()
	*sum += square
	m.Unlock()

	wg.Done()
}

func main() {
	sum := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 2; i <= 10; i += 2 {
		wg.Add(1)
		go sumSquaresNumbers(i, &sum, &wg, &m)
	}
	wg.Wait()

	fmt.Printf("Total sum squares of numbers %v\n", sum)
}
