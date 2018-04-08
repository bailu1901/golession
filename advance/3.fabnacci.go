package lesson

import (
	"fmt"
)

func fab(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("over")
		}
	}
}

func Fab() {
	ch := make(chan int)
	quit := make(chan int)

	go fab(ch, quit)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	quit <- 1
}
