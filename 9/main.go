package main

import (
	"fmt"
)

func sendValues(arr [5]int, c chan<- int) {
	for _, v := range arr {
		c <- v
	}

	close(c)
}

func sendSquares(req <-chan int, res chan<- int) {
	for v := range req {
		res <- v * 2
	}

	close(res)
}

func main() {

	array := [5]int{2, 4, 6, 8, 10}

	values := make(chan int)
	squares := make(chan int)

	go sendValues(array, values)
	go sendSquares(values, squares)

	for v := range squares {
		fmt.Println(v)
	}
}
