package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100       // меняет значение 0 индекса массива
	v = append(v, b) // тут создается новый блочный массив на который ссылается слайс v = []int8{100, 2, 3, 4, 5, b}
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
