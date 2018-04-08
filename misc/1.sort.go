package lesson

import (
	"fmt"
	"sort"
)

type myData struct {
	a int
	b bool
}

type myDataArray []myData

func (arr myDataArray) Len() int {
	return len(arr)
}

func (arr myDataArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr myDataArray) Less(i, j int) bool {
	if !arr[i].b {
		if !arr[j].b {
			return arr[i].a < arr[j].a
		}
		return true
	}
	return arr[i].a < arr[j].a
}

func SortTest() {

	var a = []int{2, 3, 7, 3, 2, 3, 1}

	sort.Ints(a)

	fmt.Println(a)

	var s = []string{"b", "f", "a", "aww"}
	sort.Strings(s)
	fmt.Println(s)

	temp := []myData{{a: 1, b: true}, {a: 2, b: true}, {a: 1, b: false}, {a: 4, b: false}}
	sort.Sort(myDataArray(temp))
	fmt.Println(temp)
}
