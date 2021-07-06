package task20

import "fmt"

func Start() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a") // тут создается новый участок памяти на который ссылается слайс, блочной области видимости
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)

	fmt.Print(slice) // тут slice = []string{"a", "a"}, т.к. глобальная область видимости
}
