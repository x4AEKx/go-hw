package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a") // slice = []string{"a", "a", "a"}
		slice[0] = "b"             // slice = []string{"b", "a", "a"}
		slice[1] = "b"             // slice = []string{"b", "b", "a"}
		fmt.Print(slice)           // в этой блочной области slice = []string{"b", "b", "a"}
	}(slice)
	fmt.Print(slice) // тут slice = []string{"a", "a"}, т.к. глобальная область видисомти
}
