package lesson

import (
	"encoding/json"
	"fmt"
)

type s1 struct {
	Page   int
	Fruits []string
}

func JsonTest() {
	t1 := s1{Page: 1, Fruits: []string{"a", "b"}}

	b1, _ := json.Marshal(&t1)
	fmt.Println(string(b1))

	t2 := s1{}
	if err := json.Unmarshal(b1, &t2); nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(t2)
}
