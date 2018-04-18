package test

import (
	"fmt"
	"testing"
)

type mytype struct {
	a int
	b int
}

func TestChanValue(t *testing.T) {
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
		a1 := [5]int{1, 2}

		c := make(chan []int, 1)
		a := make([]int, 10)
		a[0] = 1
		a[1] = 2
		c <- a

		temp := <-c
		a[0] = 10
		t.Log(a1, a, temp)
	}

}
