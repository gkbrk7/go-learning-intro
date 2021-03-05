package pointer_struct

import "fmt"

func Pointer2(numbers []int)  {
	numbers[0] = 1000
	fmt.Println("First element: ", numbers[0])
}