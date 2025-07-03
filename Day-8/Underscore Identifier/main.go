package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denomination must not be zero....")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Error handling in go")

	ans, _ := divide(10, 0)

	fmt.Println("Division of two number is: ", ans)
}

// In go, the underscore(_) is used as a blank identifier. It serves as a write-only variable that allows you to discard values returned by a function or to ignore specific vlaues when you are not interested in using them
