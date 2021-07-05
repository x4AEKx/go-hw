package main

import (
	"fmt"
)

func arrayValues(arr [5]int, c chan<- int) {
	for _, v := range arr {
		c <- v
	}

	close(c)
}

func squaresValues(req chan int, res chan<- int) {
	for v := range req {
		res <- v * v
	}

	close(res)
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	values := make(chan int)
	squares := make(chan int)

	go arrayValues(array, values)
	go squaresValues(values, squares)

	for v := range squares {
		fmt.Println(v)
	}
}
