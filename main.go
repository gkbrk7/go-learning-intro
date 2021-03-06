package main

import (
	// "fmt"
	// "golesson/packagesmodules"
	// "golesson/conditionals"
	// "golesson/loops"
	// "golesson/arrays"
	// "golesson/slices"
	// "golesson/functions"
	// "golesson/maps_forrange"
	// "golesson/pointer_struct"
	// "golesson/goroutines"
	// "golesson/gochannels"
	"golesson/interfaces"
	"golesson/stringmethods"
	// "time"
)

func main() {
	// packagesmodules.Hello()
	// conditionals.Conditional()
	// conditionals.Biggest()
	// loops.Loop()
	// loops.Guess()
	// loops.IsPrime()
	// arrays.Array()
	// slices.Slice()
	// functions.SayHello()
	// fmt.Println(functions.Sum(2,3))
	// fmt.Println(functions.Calculator(9,3))
	// // res, res1, res2, res3 := functions.Calculator(9,3)
	// // res, res1, res2, _ := functions.Calculator(9,3) //_ means does not take last result
	// functions.VariadicSum(1,2,3,4,5,6,7,8,9)
	// functions.VariadicSum([]int{1,2,3,4,5,6,7,8,9}...)
	// maps_forrange.Map()
	// num := 20
	// pointer_struct.Pointer(&num)
	// fmt.Println(num)
	// nums := []int{1,2,3}
	// pointer_struct.Pointer2(nums)
	// fmt.Println(nums)
	// pointer_struct.Struct()

	// go goroutines.Even()
	// go goroutines.Odd()

	// evenNumbersSumCn := make(chan int)
	// oddNumbersSumCn := make(chan int)
	// go gochannels.Even(evenNumbersSumCn)
	// go gochannels.Odd(oddNumbersSumCn)

	// evenNumbersSum, oddNumbersSum := <-evenNumbersSumCn, <-oddNumbersSumCn
	// multiply := evenNumbersSum * oddNumbersSum
	// fmt.Println("Multiply : ", multiply)
	
	// time.Sleep(time.Second * 2)
	interfaces.Test()
	interfaces.CreditTest()
	stringmethods.Test()
}
