package lesson

import (
	"fmt"
	"os"
)

func DeferTest() {
	f := createFile("./tmp")

	defer closeFile(f)

	write(f, "teafdas")
}

func createFile(filename string) *os.File {
	fmt.Println("create file")
	f, err := os.Create(filename)
	if nil != err {
		panic(err)
	}

	return f
}

func write(file *os.File, content string) {
	fmt.Println("write", content)
	fmt.Fprint(file, content)

}

func closeFile(file *os.File) {
	fmt.Println("close file")
	file.Close()
}
