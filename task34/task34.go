package task34

import (
	"fmt"
	"go-hw/task14"
	"sync"
	"unicode/utf8"
)

func isUniq(str string) bool {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	set := task14.NewSet(&mu, &wg)

	for _, letter := range str {
		wg.Add(1)
		go set.Add(string(letter))
	}

	wg.Wait()

	return set.Lenght() == utf8.RuneCountInString(str) //3
}

func Start() {
	str := "cat"

	value := isUniq(str)

	fmt.Println(value)
} // map["c": struct{}{}, "a": struct{}{}, "t": struct{}{}], 3
