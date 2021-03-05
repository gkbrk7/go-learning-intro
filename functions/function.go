package functions

import "fmt"

func Sum(num int, num2 int) int {
	return num + num2
}

func SayHello(){
	fmt.Println("Hello World!!")
}

// Multi return functions
func Calculator(num int, num2 int) (int, int, int, float32) {
	sum := num + num2
	multiply := num*num2
	subtract := num-num2
	divide := float32( num / num2)

	return sum, multiply, subtract, divide
}

// Variadic functions
func VariadicSum(nums ...int){
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}