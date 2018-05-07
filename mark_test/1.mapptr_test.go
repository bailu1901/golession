package golesson

import (
	"fmt"
	"testing"
)

type myData struct {
	a int
	b string
}

var (
	dataList []myData
	dataMap  map[int]*myData
)

func TestMapPtr(t *testing.T) {
	dataList = []myData{
		{a: 1, b: "1"},
		{a: 2, b: "2"},
	}

	dataMap = make(map[int]*myData)

	//这里只是把v临时变量的指针保存起来了
	//而且每次v都会修改，然后map的内容也会修改，最后map所有数据都一样，并且是最后一次赋值的内容
	for _, v := range dataList {
		dataMap[v.a] = &v
	}

	for k, v := range dataMap {
		fmt.Println(dataMap[k])
		fmt.Println(*v)
	}

	//map[1:0xc04205c460 2:0xc04205c460]
	//
}
