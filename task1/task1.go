package task1

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	complite bool
} // Композиция

func NewAction(name string, age int, complite bool) *Action {
	return &Action{
		Human: Human{
			name: name,
			age:  age,
		},
		complite: complite,
	}
} // Композиция

func SuperNewAction(h Human, complite bool) *Action {
	return &Action{
		Human:    h,
		complite: complite,
	}
} // Агрегация

func Start() {
	Action1 := NewAction("vasya", 10, true)

	Action2 := SuperNewAction(Human{name: "dima", age: 10}, false)

	fmt.Println(Action1)
	fmt.Println(Action2)
}
