package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-t1.C:
			fmt.Println("t1=", time.Now())
		case <-t2.C:
			fmt.Println("t2=", time.Now())

		}
	}
}
