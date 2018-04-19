package test

import (
	"testing"
)

type T struct {
	name string
}

func (t T) M1() {
	t.name = "name2"
}

func (t *T) M2() {
	t.name = "name2"
}

type Info interface {
	M1()
	M2()
}

type S struct {
	T
}

func TestPtr(_ *testing.T) {
	t1 := &T{}
	t1.M1()
	t1.M2()

	s1 := &S{}
	s1.M1()
	s1.M2()

	t2 := T{}
	var i Info = &t2
	i.M1()
	i.M2()
}
