package main

import "fmt"

func simplFunction() {

}
func add(a, b int) int {
	return a + b
}

//function with named return variable
func multiply(x, y int) (result int) {
	result = x * y
	return
}
func main() {
	fmt.Println("we ar learing function in go lang")

	ans := add(3, 4)
	fmt.Println(ans)
	simplFunction()

	result := multiply(4, 5)
	fmt.Println(result)
}
