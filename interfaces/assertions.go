package interfaces

import "fmt"

func Verify(i interface{})  {
	number,ok := i.(int)
	fmt.Println(number, ok)
}

func Assertion()  {
	var num interface{} = "Gokberk"
	Verify(num)
	
	var _num interface{} = 5
	Verify(_num)
}