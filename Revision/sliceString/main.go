package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"apple", "mango", "orange"}

	for _, val := range data {
		firstNames := strings.ToLower(string(val[0]))
		if strings.Contains("aieou", firstNames) {
			fmt.Println(val)
		}
	}
}
