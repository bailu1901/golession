package lesson

import (
	"fmt"
)

func SliceTest() {

	var s1 = make([]int, 3)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3

	for k, v := range s1 {
		s1[k] = v + 1
	}
	fmt.Println(s1)

	s2 := append(s1, 10)
	for k, v := range s2 {
		s2[k] = v + 1
	}
	fmt.Println(s1[:2])
	fmt.Println(s2[1:])

	copy(s2, s1)
	fmt.Println(s2)

	l := s2[:3]
	l[0] = 0
	fmt.Println(s2)
}
