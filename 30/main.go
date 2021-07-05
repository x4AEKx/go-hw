package main

import "fmt"

func main() {
	a := []int{1, 6, 7, 5}
	i := 1

	fmt.Println(a)

	a = append(a[:i], a[i+1:]...) // сливает срез и значения, срез от начальной позиции до i, не включительно; значения от a[i+1:]... - разворачивает значения

	fmt.Println(a)
}
