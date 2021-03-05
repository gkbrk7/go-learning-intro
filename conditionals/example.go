package conditionals

import "fmt"

func Biggest(){
	var number1,number2,number3 int = 10, 73, 5
	var max int = number1

	if number2 > max {
		max = number2
	}

	if number3 > max {
		max = number3
	}

	fmt.Printf("\nMaximum Value : %v\n", max)
}