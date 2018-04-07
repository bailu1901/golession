package lesson

import(
	"fmt"
	
)

func cacal(nums... int)int{
	c:=0
	for _,v := range nums{
		c+=v
	}
	return c
}

func VaradicTest(){
	arr := []int{1,2,3,4}
	fmt.Println( cacal(arr...))
}