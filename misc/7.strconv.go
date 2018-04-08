package lesson

import (
	"fmt"
	"strconv"
)

func StrConvTest() {
	p := fmt.Println

	f, _ := strconv.ParseFloat("1.2343", 64)
	p(f)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
}
