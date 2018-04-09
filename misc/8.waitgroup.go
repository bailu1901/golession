package lesson

import (
	"sync"
)

func WaitGroupTest() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ii int) {
			println(ii)
			wg.Done()
		}(i)
	}
	println("begin")
	wg.Wait()
	println("end")
}
