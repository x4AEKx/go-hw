package main

import "fmt"

func arrayValues(arr [5]int, c chan int) {
	for _, v := range arr {
		c <- v
	}

	close(c)
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	channel := make(chan int)

	go arrayValues(array, channel)

	for v := range channel {
		fmt.Println(v)
	}
}
