package main

import "fmt"

type Course struct {
	title   string
	modules []string
	rating  map[string]int
}

func main() {
	c := Course{
		title:   "Go programming",
		modules: []string{"Basics", "struct", "variable", "map"},
		rating: map[string]int{
			"bob":  5,
			"john": 4,
		},
	}
	fmt.Println("Course Title: ", c.title)
	fmt.Println("Modules")
	for i, m := range c.modules {
		fmt.Printf(" %d. %s\n", i+1, m)
	}
	fmt.Println("rating")
	for user, rate := range c.rating {
		fmt.Printf(" %s rated: %d\n", user, rate)
	}
}
