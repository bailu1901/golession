package test

import "testing"

func TestDefer(t *testing.T) {
	a := 10

	f := func(temp int) {
		t.Error(temp)
	}

	defer f(a)

	a = 20
}
