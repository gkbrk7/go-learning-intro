package interfaces

import (
	"fmt")


type Mortgage struct {
	creditPaymentTotal float64
	address string
	rate float64
}

type Car struct{
	creditPaymentTotal float64
	carInfo string
	rate float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator)  {
	paymentTotal := 0.0
	for _, credit := range credits {
		paymentTotal += credit.Calculate()
	}
	fmt.Println(paymentTotal)
}

func CreditTest()  {
	credits1 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credits2 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credits3 := Car{rate: 15, creditPaymentTotal : 60000, carInfo: "BMW"}
	
	credits := []CreditCalculator{credits1, credits2, credits3}
	CalculateMonthlyPayment(credits)
}