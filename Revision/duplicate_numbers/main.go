package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 1, 3, 5, 6}
	seen := make(map[int]int)
	duplicates := []int{}

	// Count each number
	for _, i := range num {
		seen[i]++
	}

	// Collect duplicates (count > 1)
	for value, count := range seen {
		if count > 1 {
			duplicates = append(duplicates, value)
		}
	}

	fmt.Println("Duplicates:", duplicates)
}
