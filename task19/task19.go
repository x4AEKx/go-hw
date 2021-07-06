package task19

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "Д"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10) // двоичный сдвиг 1 << 10 = (1 * 2 - 10 раз) = 1024

	v1 := []rune(v)
	justString = string(v1[:100])
}

func Start() {
	word1 := "Vova"
	word2 := []rune("Вова")

	fmt.Println(word1[:3])
	fmt.Println(string(word2[:3]))

	someFunc()
}
