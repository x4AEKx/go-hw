package task24

import "fmt"

func Start() {
	slice := make([]int, 0, 100)
	fmt.Printf("slice = %v; len(slice) = %v; cap(slice) = %v\n", slice, len(slice), cap(slice))
}
