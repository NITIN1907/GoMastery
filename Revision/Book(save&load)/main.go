package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	BookId int    `json:"book_id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	book := Book{01, "go lang developer", "Nitin"}
	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error marshalling...")
	}
	os.WriteFile("book.json", data, 0644)

	readData, _ := os.ReadFile("book.json")
	var load Book
	json.Unmarshal(readData, &load)
	fmt.Println("load file: ", load)
}
