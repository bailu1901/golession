package lesson

import (
	"fmt"
)

func IfTest() {
	var i = 0
	if 1 == i {
		fmt.Println("1")
	} else if i += 1; 1 == i {
		fmt.Println("2")
	} else if i += 1; 3 == i {
		fmt.Println("3")
	}
}
