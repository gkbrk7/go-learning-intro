package main

import "fmt"

func main() {
	var text string = "Hello World"
	var number int = 12
	var number2 float32 = 10.2
	var boolean bool = true
	number3 := 40
	number = 25
	fmt.Println(text + " ")
	fmt.Println(boolean)
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Printf("Data type: %T\n", number3)
}
