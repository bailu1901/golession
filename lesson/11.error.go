package lesson

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 0 {
		msg := fmt.Sprintf("%d can't work", arg)
		return 0, errors.New(msg)
	}
	return 100 / arg, nil
}

type argError struct {
	num int
}

func (e *argError) Error() string {
	return fmt.Sprintf("argError: %d can't work", e.num)
}

func f2(arg int) (int, error) {
	if 0 == arg {
		return 0, &argError{num: arg}
	}
	return 100 / arg, nil
}

func ErrorTest() {
	_, e := f1(0)
	if nil != e {
		fmt.Println(e.Error())
	}

	_, e = f2(0)
	if nil != e {
		fmt.Println(e.Error())
	}
}
