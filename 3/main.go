package main

import (
	"fmt"
)

func generateNumbers(c chan<- int) {
	for i := 2; i <= 10; i += 2 {
		c <- i
	}

	close(c)
}

func squaresNumbers(req <-chan int, res chan<- int) {
	for v := range req {
		res <- v * v
	}

	close(res)
}

func sumSquaresNumbers(req <-chan int, res chan<- int) {
	var total int

	for v := range req {
		total += v
	}

	res <- total
	// Где-то прочитал, каналы не файлы их можно не закрывать, например когда вывожу 1 сообщение
}

func main() {

	numbers := make(chan int)
	squares := make(chan int)
	result := make(chan int)

	go generateNumbers(numbers)
	go squaresNumbers(numbers, squares)
	go sumSquaresNumbers(squares, result)

	fmt.Printf("Total sum squares of numbers %v\n", <-result)
}
