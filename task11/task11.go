package task11

import "fmt"

func Start() {
	unorderedArray := [7]int{2, 4, 6, 8, 10, 0, 12}
	unorderedArray2 := [9]int{1, 3, 5, 0, 2, 7, 9, 11, 13}

	for _, v := range unorderedArray {
		for _, v2 := range unorderedArray2 {
			if v == v2 {
				fmt.Println(v)
			}
		}
	}

}
