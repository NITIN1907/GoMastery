package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[1:4] // slice from index 1 to 3 (1 to n-1)

	fmt.Println("Slice: ", slice)
	fmt.Println("Length(number of elements in the slice): ", len(slice))
	fmt.Println("Capacity(number of elements in the underlying array starting from the index of the slice): ", cap(slice))

}
