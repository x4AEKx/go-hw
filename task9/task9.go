package task9

import (
	"fmt"
	"time"
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

func print(req <-chan int, res chan<- struct{}) {
	for v := range req {
		time.Sleep(time.Second)
		fmt.Println(v)
	}

	res <- struct{}{}
}

func Start() {

	array := [5]int{2, 4, 6, 8, 10}

	values := make(chan int)
	squares := make(chan int)
	done := make(chan struct{})

	go sendValues(array, values)
	go sendSquares(values, squares)
	go print(squares, done)

	<-done
}
