package defers

import "fmt"

func F(number int32) string {
	if number % 2 == 0 {
		return "Even Number"
	}

	if number % 2 != 0{
		return "Odd Number"
	}

	return "Unknown"
}

func Testing()  {
	defer DeferFunc()
	fmt.Println(F(1))
}

func DeferFunc()  {
	fmt.Println("Finished")
}