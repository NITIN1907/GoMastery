package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
	fmt.Println(*num)
}
func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	fmt.Println("num has value: ", num)
	fmt.Println("pointer's value: ", ptr)
	fmt.Println("value in the address: ", *ptr)

	var pointer *int // default value of pointer is nil

	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}
	value := 20
	modifyValueByReference(&value)
	fmt.Println("Value contains: ", value)

}
