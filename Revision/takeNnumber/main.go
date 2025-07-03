package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the numbers of element: \n")
	fmt.Scan(&n)
	var count int
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("enter the numbers:\n")
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			count++
		}
		if arr[i]%2 == 0 {
			fmt.Println("number of even number: ", arr[i])
		}
	}
	fmt.Print("Count of negative numbers: ", count)
}
