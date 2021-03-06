package exceptionhandling

import (
	"fmt"
	"os"
)

func Exception()  {
	f,err := os.Open("demo1.txt")
	// If file exists err assigned as nil
	if err != nil {
		// Type assertion
		if pErr,ok := err.(*os.PathError); ok{
			fmt.Println("File does not exist: ", pErr.Path)
			return
		}else{
			fmt.Println("File does not exist")
			return
		}
	}else{
		fmt.Println(f.Name())
	}
}