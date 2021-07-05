package main

import "fmt"

type set map[string]struct{}

func NewSet() set {
	return make(map[string]struct{})
}

func (s set) Add(str string) {
	_, ok := s[str]
	if ok {
		return
	} else {
		s[str] = struct{}{}
	}
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	mySet := NewSet()

	for _, v := range strings {
		mySet.Add(v)
	}

	fmt.Println(mySet)
}
