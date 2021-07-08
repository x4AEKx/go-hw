package task36

import "fmt"

// Дан массив из 10 строк, посчитать в каждой строке количество уникальных символов с помощью горутин.

func arrToStr(arr [10]string, c chan string) {
	for _, v := range arr {
		c <- v
	}
	close(c)
}

func uniqCount(req <-chan string, res chan<- int) {

	for str := range req {
		m := make(map[rune]bool)
		count := 0

		for _, v := range str {
			_, ok := m[v]
			if ok {
				m[v] = false
			} else {
				m[v] = true
			}
		}

		for _, value := range m {
			if value {
				count++
			}
		}

		res <- count
	}
	close(res)
}

func Start() {
	arr := [10]string{"мам", "dad", "cat", "dog", "tree", "cda", "aaa", "cda", "adf", "afg"}
	// должно быть = 1, 1, 3, 3, 2, 3, 0, 3, 3, 3

	strings := make(chan string)
	count := make(chan int)

	go arrToStr(arr, strings)
	go uniqCount(strings, count)

	for v := range count {
		fmt.Println(v)
	}
}
