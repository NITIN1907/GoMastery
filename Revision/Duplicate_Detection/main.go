package main

import "fmt"

func main() {
	names := []string{"apple", "banana", "orange", "apple", "orange"}

	seen := make(map[string]bool)

	for _, name := range names {
		if seen[name] {
			fmt.Println(name)
		} else {
			seen[name] = true
		}
	}
}
