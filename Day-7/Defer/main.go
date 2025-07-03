package main

import "fmt"

func main() {
	//*In Go, the defer keyword is used to postpone the execution of a function until the surrounding function returns.
	fmt.Println("Starting of the program")
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
	
}
