package lesson

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var dict = map[string]int{"trest": 1}

var mut = &sync.Mutex{}

func MutexTest() {
	dict["test"] = 1

	for i := 0; i < 5; i++ {
		go func() {
			idx := 0
			for ; ; idx++ {
				mut.Lock()
				dict[fmt.Sprintf("%d", idx)] = idx
				mut.Unlock()
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Millisecond)
	mut.Lock()
	fmt.Println(dict)
	mut.Unlock()
}
