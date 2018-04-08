package lesson

import (
	"fmt"
	"time"
)

func work(id int, jobs chan int, result chan int) {
	for j := range jobs {
		fmt.Println(id, "work ret=", j)
		time.Sleep(time.Second)
		result <- j
	}
}

func WorkTest() {
	job := make(chan int, 100)
	result := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go work(i, job, result)
	}

	for i := 0; i < 10; i++ {
		job <- i
	}
	close(job)
	//time.Sleep(time.Second*10)
	for i := 0; i < 10; i++ {
		<-result
	}

}
