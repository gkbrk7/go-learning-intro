package pointer_struct

import "fmt"

func Struct()  {
	fmt.Println(product{"Computer",5000, "XYZ"})
	fmt.Println(product{name : "Computer", price : 5000})
	p := product{name : "Pineapple", price : 25, brand: "Excellence"}
	p.save()
}

type product struct {
	name string
	price float64
	brand string
}

func (p product) save()  {
	fmt.Println("Product saved")
}
