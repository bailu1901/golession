package lesson

import (
	"fmt"
	"time"
)

func LimiterTest() {

	q1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		q1 <- i
	}

	close(q1)

	t1 := time.Tick(time.Millisecond * 200)
	for q := range q1 {
		<-t1
		fmt.Println("q1", q, time.Now())
	}

	t2 := make(chan time.Time, 100)
	go func() {
		t2 <- time.Now()
		t2 <- time.Now()
		t2 <- time.Now()

		for temp := range time.Tick(time.Millisecond * 200) {
			t2 <- temp
		}
	}()

	q2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		q2 <- i
	}
	close(q2)

	for q := range q2 {
		<-t2
		fmt.Println("q2", q, time.Now())
	}
}
