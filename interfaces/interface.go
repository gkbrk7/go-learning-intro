package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r circle) area() float64 {
	return r.radius * r.radius * math.Pi
}

func testEnv(s shape)  {
	fmt.Println(s.area())
}

func Test()  {
	r:= rectangle{width : 50, height : 100}
	testEnv(r)

	c:= circle{radius: 10}
	testEnv(c)
}

