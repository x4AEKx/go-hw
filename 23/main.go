package main

import (
	"fmt"
	"sort"
)

func main() {
	sortSlice := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5}
	x := 36

	i := sort.Search(len(sortSlice), func(i int) bool { return sortSlice[i] <= x })
	fmt.Printf("Found %d at index %d in %v\n", x, i, sortSlice)
}
