package main

import "fmt"

func main() {
	slice := make([]int, 0, 100)
	fmt.Printf("slice = %v; len(slice) = %v; cap(slice) = %v\n", slice, len(slice), cap(slice))
}
