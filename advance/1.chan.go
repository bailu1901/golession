package lesson

import (
	"fmt"
)

func inChan(ch chan<- string, str string) {

	ch <- str
	//<-ch error
}

func outChan(ch <-chan string) string {
	//ch <- "a" error
	return <-ch
}

func ChanTest() {
	message := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {

			//time.Sleep(time.Second * 2)
			message <- i
			fmt.Printf("message <- %d\n", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ret := <-message
			fmt.Printf("  %d<-message\n", ret)
		}
	}()

	var input = ""
	fmt.Scanln(&input)
	//fmt.Println(ret)
}
