package lesson

import (
	"fmt"
	"time"
)

// TimeoutTest test
// @param p1 - p1
// @param p2 - p2
func TimeoutTest(p1 int, p2 int) {

	c := make(chan int)

	go func() {
		select {
		case <-c:
			fmt.Println("c")
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}()

	time.Sleep(1000)
	c <- 1
	fmt.Scanln("")

}
