package task18

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100       // меняет значение 0 индекса массива
	v = append(v, b) // тут создается новый участок в памяти на который ссылается слайс
}

func Start() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
