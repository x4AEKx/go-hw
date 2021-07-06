package task32

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
	fmt.Println("Sleep seconds: ", sec)
}

func Start() {
	fmt.Println("Start main")
	sleep(1)
	fmt.Println("Down main")
}
