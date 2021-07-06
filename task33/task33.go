package task33

import (
	"fmt"
	"math/rand"
	"time"
)

func createRndNum(c chan<- int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 10; i++ {
		c <- rand.Intn(100)
	}

	close(c)
}

func evenNum(req <-chan int, res chan<- int) {
	for v := range req {
		if v%2 == 0 {
			res <- v
		}
	}

	close(res)
}

func Start() {
	randNumbers := make(chan int)
	even := make(chan int)

	go createRndNum(randNumbers)
	go evenNum(randNumbers, even)

	for v := range even {
		fmt.Println(v)
	}
}
