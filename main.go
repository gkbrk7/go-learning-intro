package main

import (
	"fmt"
	"golesson/packagesmodules"
	"golesson/conditionals"
	"golesson/loops"
	"golesson/arrays"
	"golesson/slices"
	"golesson/functions"
)

func main() {
	packagesmodules.Hello()
	conditionals.Conditional()
	conditionals.Biggest()
	loops.Loop()
	loops.Guess()
	loops.IsPrime()
	arrays.Array()
	slices.Slice()
	functions.SayHello()
	fmt.Println(functions.Sum(2,3))
}
