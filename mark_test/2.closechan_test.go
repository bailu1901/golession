package golesson

import (
	"fmt"
	"sync"
	"testing"
)

func TestCloseChan(t *testing.T) {

	c := make(chan *int, 1)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				select {
				case ret := <-c:
					fmt.Println(idx, ret)
					return
				}
			}

		}(i)
	}

	//意义在于当一个chan被关闭了，所有取通道值得地方都会得到通知，并在通道中得到一个默认值，常用于退出状态
	close(c)

	wg.Wait()

}
