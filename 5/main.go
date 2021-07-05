package main

import (
	"fmt"
	"time"
)

func write(c chan<- int) {
	i := 0

	for {
		c <- i
		i++
	}

}

func read(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	channel := make(chan int)

	go write(channel)
	go read(channel)

	time.Sleep(time.Second)
	close(channel)
}
