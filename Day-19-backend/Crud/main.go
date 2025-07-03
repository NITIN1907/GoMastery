package main

import (
	"encoding/json"
	"fmt"

	// "io"
	"net/http"
)

type Todo struct {
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

	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("error in getting Reponse...", res.Status)
	// 	return
	// }
	// data, errors := io.ReadAll(res.Body)
	// if errors != nil {
	// 	fmt.Println("Error reading...", err)
	// 	return
	// }

	// fmt.Println(string(data))

	var todo Todo
	er := json.NewDecoder(res.Body).Decode(&todo)
	if er != nil {
		fmt.Println("Error in decoding: ", er)
	}
	fmt.Println(todo)

}
