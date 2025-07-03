package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
	customer
}

func main() {
	// newCust := customer{
	// 	name:  "abc",
	// 	phone: "95958585858",
	// }
	newOrder := order{
		id:     "1",
		amount: 30.0,
		status: "recieved",
		// customer: newCust,
		customer: customer{
			name:  "abc",
			phone: "4594904094",
		},
	}

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}
