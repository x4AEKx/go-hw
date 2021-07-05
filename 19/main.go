package main

import (
	"fmt"
	"unsafe"
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
	justString = v[:100]

	fmt.Println("cap v =", cap([]byte(v)), "len v =", len(v), "sizeof(v) =", unsafe.Sizeof(v), v)

	fmt.Println("cap justString =", cap([]byte(justString)), "len justString =", len(justString), "sizeof(justString) =", unsafe.Sizeof(justString), justString)
}

func main() {
	someFunc()
}
