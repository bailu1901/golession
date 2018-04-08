package lesson

import (
	"fmt"
)

func indexOf(vs []string, s string) int {
	for i, v := range vs {
		if v == s {
			return i
		}
	}
	return -1
}

func filter(vs []string, f func(string) bool) []string {
	tmp := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func makeMap(vs []string, f func(string) bool) map[int]string {
	tmp := make(map[int]string)
	for i, v := range vs {
		if f(v) {
			tmp[i] = v
		}

	}
	return tmp
}

func FuncTest() {
	var temp = []string{"a", "b", "c"}

	fmt.Println(filter(temp, func(v string) bool {
		return v == "a"
	}))
}
