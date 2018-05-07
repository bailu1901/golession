package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type mytype struct {
	a int
	b int
}

func TestChanValue(t *testing.T) {

	/*
		{
			c := make(chan *mytype, 1)

			a := &mytype{a: 1, b: 2}
			c <- a
			temp := <-c
			temp.a = 10
			fmt.Println(a, temp)
			t.Log(*a, *temp)
		}

		{
			a := [5]int{1, 2}

			c := make(chan [5]int, 1)
			c <- a
			a1 := <-c

			a1[0] = 2

			fmt.Println(a, a1)
		}

		{
			c := make(chan int)
			go func() {
				i := 0
				ticker := time.NewTicker(time.Second)
				for {
					select {
					case <-ticker.C:
						c <- i
					}
					i++
					if 10 == i {
						close(c)
					}
				}

			}()

			for r := range c {
				fmt.Println(r)
			}
		}
	*/

	{
		closeChan := make(chan bool)

		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(idx int) {

				for {
					select {
					case v := <-closeChan:
						fmt.Println("over", idx, v)
						wg.Done()
						return
					}
				}
			}(i)

		}

		closeChan <- true
		time.Sleep(time.Second)
		close(closeChan)

		wg.Wait()
	}

	{ /*
			closeChan := make(chan int, 10)

			go func() {

				for i := 0; i < 10; i++ {
					closeChan <- i

					time.Sleep(time.Millisecond * 200)
				}
				close(closeChan)
			}()

			for v := range closeChan {
				fmt.Println(v)
			}
		*/
	}

}
