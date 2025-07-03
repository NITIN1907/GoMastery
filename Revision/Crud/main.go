package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type placeholder struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning crud...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting.....", err)
		return
	}

	defer res.Body.Close()

	data, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error reading...", err)
		return
	}

	fmt.Println(string(data))

	var place placeholder
	er := json.Unmarshal(data, &place)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Printf("Parsed Struct: %+v\n", place)
}
