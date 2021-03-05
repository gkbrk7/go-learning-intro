package loops

import "fmt"

func Guess(){
	Guess := 0
	Number := 80

	for i := 0; i < 5; i++ {
		fmt.Scanln(&Guess)
		if Guess == Number {
			fmt.Println("Your Guess is Correct")
			break
		}else {
			fmt.Printf("Your Guess is not Correct: Right : %v\n", 5 - (i+1))
			continue
		}
	}
}