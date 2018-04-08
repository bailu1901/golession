package lesson

import (
	"fmt"
)

func RangeChanTest() {
	c := make(chan int, 10)

	c <- 10
	c <- 20

	close(c)

	for elem := range c {
		fmt.Println(elem)
	}
}
