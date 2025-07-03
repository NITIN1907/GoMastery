package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of student...\n")
	fmt.Scan(&n)

	marks := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter marks of student: %d\n", i+1)
		fmt.Scan(&marks[i])
	}

	var sum, max, min float64
	max = marks[0]
	min = marks[0]
	for i := 0; i < n; i++ {
		sum += marks[i]
		if max < marks[i] {
			max = marks[i]
		}
		if min > marks[i] {
			min = marks[i]
		}
	}

	avg := sum / float64(n)

	fmt.Println("Average: ", avg)
	fmt.Println("Highest: ", max)
	fmt.Println("lowest: ", min)

}
