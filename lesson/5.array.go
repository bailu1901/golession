package lesson

import (
	"fmt"
)

func ArrayTest() {

	var a1 [5]int
	a1[0] = 1

	a2 := [5]int{10, 10}
	fmt.Println(a2)

	var a3 [2][2]float32
	a3[0][0] = 0
	//append(a,1)
}
