package task26

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Start() {
	str := "string строка"

	fmt.Println("Started string = ", str)

	revStr := reverse(str)

	fmt.Println("Finished string = ", revStr)
}
