package main

import "fmt"

func sum(nums ...int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	return total
}

func anyType(res ...any) {
	for _, result := range res {
		fmt.Println(result)
	}
}
func main() {
	fmt.Println("variadic function")
	// result := sum(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)

	anyType(1, 2, 3, "anytype", "hello")

}
