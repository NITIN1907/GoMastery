package main

import "fmt"

func counter() func() int {
	var counter int
	return func() int {
		counter += 1
		return counter
	}
}
func main() {
	fmt.Println()
	makeCounter := counter()
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())

}
