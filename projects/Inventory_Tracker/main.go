package main

import "fmt"

type Item struct {
	name     string
	quantity int
	price    float64
}

func addStock(item *Item, qty int) {
	item.quantity += qty
	fmt.Println("Item added...")
}
func updatePrice(item *Item, newPrice float64) {
	item.price = newPrice
	fmt.Println("Item price updated..")
}
func removeStock(item *Item, qty int) {
	if item.quantity >= qty {
		item.quantity -= qty
		fmt.Println("item removed....")
	} else {
		fmt.Printf("Can't remove %d %s(s). Only %d in stock.\n ", qty, item.name, item.quantity)
	}
}
func printInventory(items []*Item) {
	fmt.Println("\n-----Inventory-------")
	for _, item := range items {
		fmt.Printf("Name: %-10s | Quantity: %-3d | Price: $%.2f\n", item.name, item.quantity, item.price)

	}
	fmt.Println("--------------------")
}
func main() {
	it1 := &Item{name: "pencil", quantity: 10, price: 10.5}
	it2 := &Item{name: "paper", quantity: 20, price: 20.0}
	it3 := &Item{name: "scissor", quantity: 10, price: 30.0}

	inventory := []*Item{it1, it2, it3}

	printInventory(inventory)

	addStock(it1, 20)
	addStock(it2, 5)

	printInventory(inventory)

	removeStock(it1, 10)
	updatePrice(it3, 50.0)
	printInventory(inventory)

}
