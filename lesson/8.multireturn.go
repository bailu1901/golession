package lesson

import(
	"fmt"
	
)

func test()(int,string){
	return 1,"a"
}

func MultiReturnTest(){
	fmt.Println(test())
}
