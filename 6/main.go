package main

import (
	"fmt"
	"time"
)

func f(flag *bool) {
	for { // бесконечный цикл for опрашивает флаг
		if *flag { // если флаг = true, бесконечно выводится "gorutine working"
			fmt.Println("gorutine working")
		} else { // есди флаг = false, водится сообщение "gorutine done" и цикл for прерывается
			fmt.Println("gorutine done")
			return
		}
	}

}

func main() {
	flag := true
	go f(&flag) // 2. Горутина f выводит "gorutine working"

	time.Sleep(time.Second) // 1. Sleep заблокировал main рутину на 1 секунду, выполнение перешло к п.2

	flag = false // 3. Через 1 сек main рутина разблокировалась, флаг поменял значение на false

	time.Sleep(time.Second)
	fmt.Println("main done")
}
