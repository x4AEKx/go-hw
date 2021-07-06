package task8

import "fmt"

func Start() {
	var value int64 = 5 // 5/2 (1), 2/2 (0), (1), = 101; 7/2 (1), 3/2 (1), (1) = 111
	var i int64 = 1

	fmt.Printf("Old value = %v\n", value)

	value = value | 1<<i // устанавливает i бит в 1

	fmt.Printf("New value = %v\n", value)

	// побитовая логика го
}
