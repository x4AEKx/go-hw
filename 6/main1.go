package main

import (
	"fmt"
	"time"
)

func f(stop chan bool) {
	for {
		select { // select опрашивает была ли отправка в канал stop
		case <-stop:
			fmt.Println("gorutine done")
			return
		default: // пока отправка в канал stop не произошла и читать из канала нечего, выполняется default
			fmt.Println("gorutine working")
		}
	}
}

func main() {

	stop := make(chan bool)
	go f(stop) // 2. Горутина f выводит "gorutine working"

	time.Sleep(time.Second) // 1. Sleep заблокировал main рутину на 1 секунду, выполнение перешло к п.2
	stop <- true            // 3. Через 1 сек main рутина разблокировалась, в канал отправилось значение

	time.Sleep(time.Second)
	fmt.Println("main done")

}
