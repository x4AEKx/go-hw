package task10

import "fmt"

func group(numbers ...float32) map[int][]float32 {
	myMap := make(map[int][]float32)

	for _, num := range numbers {
		key := int(num/10) * 10
		myMap[key] = append(myMap[key], num)
	}
	return myMap
}

func Start() {
	myGroup := group(32.5, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0)

	fmt.Println(myGroup)
}
