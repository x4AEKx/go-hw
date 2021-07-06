package main

import (
	"fmt"
)

func main() {
	type Human struct {
		name string
	}

	type Action struct {
		Human
	}

	Dima := &Action{}
	Dima.name = "Dima"

	fmt.Println(Dima.name)
}
