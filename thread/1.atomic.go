package lesson

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var c uint64 = 0

func AutomicTest() {
	call := func() {
		for {
			atomic.AddUint64(&c, 1)
			runtime.Gosched()
		}
	}

	for i := 0; i < 50; i++ {
		go call()

	}

	time.Sleep(time.Second * 2)

	ret := atomic.LoadUint64(&c)
	fmt.Println(ret)
}
