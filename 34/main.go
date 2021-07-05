package main

import "fmt"

func isUniq(str string) bool {
	runes := []rune(str)

	set := make(map[rune]struct{})

	for _, v := range runes {
		_, ok := set[v]
		if ok {
			return false
		} else {
			set[v] = struct{}{}
		}
	}

	return true
}

func main() {
	str := "cat"

	value := isUniq(str)

	fmt.Println(value)
}
