package main

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100
	debitAmt := func(debit int) int {
		amount -= debit
		return amount
	}
	return debitAmt
}
func active() func(int) int {
	amount := 100
	credit := func(cred int) int {
		amount += cred
		return amount
	}
	return credit
}
func counter() func() int {
	var counter int
	fmt.Println(counter)
	return func() int {
		counter += 1
		return counter
	}
}
func main() {
	fmt.Println("closure")
	gift := activateGiftCard()
	gift2 := activateGiftCard()
	fmt.Println(gift(10))
	fmt.Println(gift2(20))

	GetAmt := active()

	fmt.Println(GetAmt(200))
	cnt := counter()
	fmt.Println(cnt())
}
