package main

import "fmt"

func main() {
	country := make(map[string]string)

	var name string
	fmt.Scan(&name)
	var capital string
	fmt.Scan(&capital)

	country[name] = capital
	fmt.Print(country)

	var query string
	fmt.Scan(&query)

	result, exist := country[query]
	if exist {
		fmt.Print("Found: ", result)
	} else {
		fmt.Print("Not found")
	}

}
