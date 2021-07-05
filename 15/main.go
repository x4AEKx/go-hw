package main

import "fmt"

func main() {
	a := 1
	b := 2

	b, a = a, b

	fmt.Println(a, b)
}
