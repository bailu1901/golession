package lesson

import (
	"fmt"
	"time"
)

func TimeTest() {
	p := fmt.Println

	then := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Weekday())

	p(now.Location())

	diff := now.Sub(then)

	p(diff.Hours())

	p(now.Unix())
}
