package lesson

import (
	"fmt"
	"time"
)

func TimeTest() {

	fmt.Println(time.Now())
	t1 := time.NewTimer(time.Second * 2)

	<-t1.C
	fmt.Println("t1", time.Now())

	t2 := time.NewTimer(time.Second * 2)
	go func() {
		<-t2.C
		fmt.Println("t2")
	}()

	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("stop2=", stop2)
	}

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
	fmt.Println("ticker.Stop()")
	time.Sleep(time.Second * 5)
}
