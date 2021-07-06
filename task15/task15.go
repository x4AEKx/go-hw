package task15

import "fmt"

func Start() {
	a := 1
	b := 2

	b, a = a, b

	fmt.Println(a, b)
}
