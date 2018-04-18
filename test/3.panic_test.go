package test

import (
	"fmt"
	"testing"
)

func recoverTest() {
	ret := recover()
	fmt.Println(ret)
}

func TestPanic(t *testing.T) {
	defer func() {
		recoverTest()
	}()

	panic("tes")

	t.Error("over")
}
