package lesson

import (
	"fmt"
)

func ForTest(){
	for  i :=0; i<3; i++{
		fmt.Println(i)
	}

	var arr = [5]int{1,2,3,4,5}
	for k,v :=range arr{
		fmt.Println(k,v)
	}
}