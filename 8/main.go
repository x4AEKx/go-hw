package main

import "fmt"

func main() {
	var value int64 = 1
	var i int64 = 1

	fmt.Printf("Old value = %v\n", value)

	value = value | 1<<i // устанавливает i бит в 1

	fmt.Printf("New value = %v\n", value)
}
