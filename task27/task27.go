package task27

import (
	"fmt"
	"strings"
)

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func reverseWords(s string) string {
	slice := strings.Split(s, " ")

	revers := reverse(slice)

	result := strings.Join(revers, " ")
	return result
}

func Start() {
	words := "snow dog sun"

	fmt.Println("Started string1 = ", words)

	reverse := reverseWords(words)

	fmt.Println("Finished string1 = ", reverse)
}
