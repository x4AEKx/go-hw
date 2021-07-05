package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(c <-chan interface{}) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	count := 10

	channel := make(chan interface{})
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < count; i++ {
		go worker(channel)
	}

	for {
		select {
		case <-sigs:
			fmt.Println("main done")
			return
		case channel <- 1:
		case channel <- "string":
		case channel <- true:
		}
	}

}
