package lesson

import (
	"fmt"
)

func MapTest() {

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m)

	fmt.Println("len=", len(m))
	delete(m, "a")

	if v, has := m["a"]; !has {
		fmt.Println(v, has)
	}

	t := m["b"]
	fmt.Println(t)

	fmt.Println(m)

	m1 := map[int]string{1: "1", 2: "2"}
	fmt.Println(m1)
}
