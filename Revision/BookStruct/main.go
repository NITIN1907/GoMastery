package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b Book) DiscountedPrice(percent float64) float64 {
	discount := (b.Price * percent) / 100
	return b.Price - discount
}

func (b Book) Display(percent float64) {
	fmt.Println("Book Title: ", b.Title)
	fmt.Println("Book Author: ", b.Author)
	fmt.Println("Book Price: ", b.Price)
	fmt.Println("Price after Discount:", b.DiscountedPrice(percent))
}

func main() {
	b := Book{
		Title:  "Go Language",
		Author: "Nitin",
		Price:  250.00,
	}
	b.Display(20.0)
}
