package loops

import "fmt"

func IsPrime(){
	p:= 0
	var flag bool = true
	fmt.Println("Give a number to check primality : ")
	fmt.Scanln(&p)

	for i := 2; i < p; i+=2 {
		if p % i == 0 {
			flag = false
			fmt.Println("Given number is not prime")
			break
		}
	}

	if flag {
		fmt.Println("Given number is prime")
	}
}