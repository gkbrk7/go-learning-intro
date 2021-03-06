package gochannels

import "fmt"

func Even(eChan chan int)  {
	sum := 0
	for i := 0; i <= 10; i+=2 {
		sum += i
	}
	eChan <- sum
}

func Odd(oChan chan int)  {
	sum := 0
	for i := 1; i <= 10; i+=2 {
		sum += i
	}
	oChan <- sum
}