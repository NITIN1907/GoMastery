package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request", err)
	}
	client := http.Client{}
	res, er := client.Do(req)
	if er != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status)

}
func main() {
	performDeleteRequest()
}
