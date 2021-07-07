package task35

import (
	"fmt"
	"strconv"
)

func compress(str string) string {
	slice := []rune(str)

	count := 0
	last := slice[0]
	mystr := ""

	for _, v := range slice {
		if v == last {
			count++
		} else {
			mystr += strconv.Itoa(count) + string(last)
			last = v
			count = 1
		}
	}

	newstr := mystr + strconv.Itoa(count) + string(last)

	return newstr
}

func Start() {
	// АААBBCCCDDEEE -> 3A2B3C2D3E

	str := "АААBBCCCDDEEE"

	compressStr := compress(str)

	fmt.Println(compressStr)

}
