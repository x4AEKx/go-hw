package task14

import (
	"fmt"
	"sync"
)

type set struct {
	m  map[string]struct{}
	mu *sync.RWMutex
	wg *sync.WaitGroup
}

func NewSet(mu *sync.RWMutex, wg *sync.WaitGroup) *set {
	return &set{
		m:  make(map[string]struct{}),
		mu: mu,
		wg: wg,
	}
}

func (s *set) Add(str string) {
	s.mu.RLock()
	_, ok := s.m[str]
	s.mu.RUnlock()
	if ok {
		s.wg.Done()
		return
	} else {
		s.mu.Lock()
		s.m[str] = struct{}{}
		s.mu.Unlock()
	}
	s.wg.Done()
}

func (s *set) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return fmt.Sprintf("%v", s.m)
}

func (s *set) Lenght() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.m)
}

func Start() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	mu := &sync.RWMutex{}

	var wg = sync.WaitGroup{}

	mySet := NewSet(mu, &wg)

	for _, v := range strings {
		wg.Add(1)
		go mySet.Add(v)
	}

	wg.Wait()
	fmt.Println(mySet)
}
