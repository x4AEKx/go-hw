package task12

import "fmt"

func update(p *int) {
	b := 2
	p = &b // чтобы исправить нужно разименовать p (*p = b)
}

func Start() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
