package defers

import "fmt"

type product struct {
	productName string
	unitPrice float64
}

func (p product) save()  {
	fmt.Println("Product Saved :", p.productName)
}

func A(){
	fmt.Println("A is running...")
}

func B(){
	// Defer function will execute after B is completely finished its process
	// Defer functions execute his process according to stack logic, All defers like stacks (LIFO)
	// Defer functions will execute if there is an error on the function
	fmt.Println("B is running...")
}

func C()  {
	defer A()
	defer B()
	fmt.Println("C is executing...")
}

func ProductTest()  {
	p := product{ productName: "Laptop", unitPrice : 5000}
	defer p.save()
	fmt.Println("Operation Successful")
}