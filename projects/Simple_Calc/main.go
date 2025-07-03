package main

import "fmt"

func main() {
	var a, b float64
	var operation string
	fmt.Println("enter the first number: ")
	fmt.Scan(&a)

	fmt.Println("enter the operator (+,-,/,x): ")
	fmt.Scan(&operation)

	fmt.Println("enter the second number: ")
	fmt.Scan(&b)

	var result float64

	if operation == "+" {
		result = a + b
	} else if operation == "-" {
		result = a - b
	} else if operation == "/" {
		result = a / b
	} else if operation == "x" {
		result = a * b
	} else {
		fmt.Println("enter the wrong operation....")
		return
	}

	fmt.Println("result: ", result)
}
