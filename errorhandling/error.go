package errorhandling

import (
	"errors"
	"fmt"
)

type borderException struct {
	message string
	parameter int
}

func Guess(g int) (string, error) {
	if g < 0 || g > 100 {
		return "",errors.New("Please give a number between 1 and 100")
	}
	return "Correct!!",nil
}

func ErrorHandler()  {
	fmt.Println(Guess(50))
	fmt.Println(Guess(101))
}

func (be *borderException) Error() string {
	return fmt.Sprintf("%d---%s", be.parameter, be.message)
}

func Guess2(g int) (string, error) {
	if g < 0 || g > 100 {
		return "", &borderException{message: "Please give a number between 1 and 100", parameter : g}
	}
	return "Correct!!",nil
}