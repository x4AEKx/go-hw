package task36

import (
	"fmt"
	"sync"
)

// Дан массив из 10 строк, посчитать в каждой строке количество уникальных символов с помощью горутин.

type set struct {
	m  map[string]bool
	mu *sync.RWMutex
	wg *sync.WaitGroup
}

func NewSet(mu *sync.RWMutex, wg *sync.WaitGroup) *set {
	return &set{
		m:  make(map[string]bool),
		mu: mu,
		wg: wg,
	}
}

func (s *set) Add(str string) {
	s.mu.RLock()
	_, ok := s.m[str]
	s.mu.RUnlock()
	if ok {
		s.mu.Lock()
		s.m[str] = false
		s.mu.Unlock()
	} else {
		s.mu.Lock()
		s.m[str] = true
		s.mu.Unlock()
	}
	s.wg.Done()
}

func (s *set) countUniq() int {
	count := 0

	for _, v := range s.m {
		if v {
			count++
		}
	}

	return count
}

func (s *set) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return fmt.Sprintf("%v", s.m)
}

func uniqCount(str string) int {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	set := NewSet(&mu, &wg)

	for _, letter := range str {
		wg.Add(1)
		go set.Add(string(letter))
	}

	wg.Wait()

	return set.countUniq()
}

func Start() {
	arr := [10]string{"мам", "dad", "cat", "dog", "tree", "cda", "aaa", "cda", "adf", "afg"}
	// должно быть = 1, 1, 3, 3, 2, 3, 0, 3, 3, 3

	for _, v := range arr {
		go fmt.Println(uniqCount(v))
	}
}
