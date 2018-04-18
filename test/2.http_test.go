package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")

	//defer resp.Body.Close() not ok

	if nil != err {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {

		fmt.Println(err)
		return
	}
	t.Error(string(body))
}
