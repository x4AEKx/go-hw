package task17

import (
	"fmt"
)

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "channel int"
	case chan string:
		return "channel string"
	case chan bool:
		return "channel bool"
	default:
		return "unknown"
	}
}

type Person struct {
	name string
}

func Start() {
	v := "hello world"
	v2 := 1
	v3 := make(chan int)
	v4 := make(chan string)
	v5 := make(chan bool)
	v6 := true
	v7 := make(chan uint)
	v8 := Person{name: "Dima"}

	fmt.Println(typeof(v))  //string
	fmt.Println(typeof(v2)) //int
	fmt.Println(typeof(v3)) //channel int
	fmt.Println(typeof(v4)) //channel string
	fmt.Println(typeof(v5)) //channel bool
	fmt.Println(typeof(v6)) //bool
	fmt.Println(typeof(v7)) //unknown
	fmt.Println(typeof(v8)) //unknown
}
