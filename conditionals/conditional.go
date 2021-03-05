package conditionals

import "fmt"

func Conditional()  {
	var account float64 = 1000
	var withdraw float64 = 900

	if withdraw > account {
		fmt.Println("Balance Not Enough")
	}else if withdraw < account {
		fmt.Printf("Total Balance: %v",account-withdraw)
	}else{
		fmt.Printf("Something went wrong!!")
	}
}