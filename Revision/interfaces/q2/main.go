package main

import "fmt"

type PaymentProcessor interface {
	pay(float64) string
}
type Stripe struct {
}
type PayPal struct {
}

func (s Stripe) pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f amount using Stripe ", amount)
}
func (p PayPal) pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f amount using PayPal ", amount)
}
func ProcessPayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.pay(amount))
}
func main() {
	s := Stripe{}
	p := PayPal{}

	ProcessPayment(s, 100)
	ProcessPayment(p, 200)
}
