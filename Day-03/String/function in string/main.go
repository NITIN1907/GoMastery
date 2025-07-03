package main

import (
	"fmt"
	"strings"
)

func main() {
	//declaring string(immutable nature)
	var st string = "Hello world, Golang"

	//functions in string
	//checking and searching
	fmt.Println("contains world:", strings.Contains(st, "world"))

	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(st, "Hello"))

	fmt.Println("Has Suffix '!': ", strings.Index(st, "o"))

	fmt.Println("LastIndex of 'o': ", strings.LastIndex(st, "o"))

	//Transforming string
	fmt.Println("To Upper: ", strings.ToUpper(st))

	fmt.Println("ToLower: ", strings.ToLower(st))

	fmt.Println("Title: ", strings.ToTitle("hello golang"))

	fmt.Println("TrimSpace:", strings.TrimSpace(" golang  "))

	fmt.Println("ReplaceAll 'o' with '@': ", strings.ReplaceAll(st, "o", "@"))

	//splitting and joining
	csv := "apple,banana,cherry"

	fruits := strings.Split(csv, ",")

	fmt.Println("Split:", fruits)

	joined := strings.Join(fruits, "+")
	fmt.Println("join: ", joined)

	sentence := "Go is fun"
	fmt.Println("Fields: ", strings.Fields(sentence))

}
