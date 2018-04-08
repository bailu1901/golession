package lesson

import (
	"fmt"
	"time"
)

var c1 = make(chan string)
var c2 = make(chan string)

func SelectTest() {
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("c1", msg1)
		case msg2 := <-c2:
			fmt.Println("c2", msg2)
		}
		//fmt.Println(<-c)
	}

}
