package lesson

import (
	"fmt"
	"time"
)

func SwitchTest() {

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			{
				fmt.Println(i)
			}
		case 1:
			{
				fmt.Println(i)
			}
		case 2:
			{
				fmt.Println(i)
			}
		}
	}

	switch time.Now().Weekday() {
	case time.Wednesday:
		fmt.Println("time.Wendsday")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("time.Hour<12")
	default:
		fmt.Println("time.Now()")
	}

}
