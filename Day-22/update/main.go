package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performUpdateRequest() {
	todo := Todo{
		UserId:    11234,
		Title:     "golang development",
		Completed: false,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling:", err)
		return
	}

	//convert string data to string
	jsonReader := strings.NewReader(string(jsonData))
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request
	req, er := http.NewRequest(http.MethodPut, myurl, jsonReader)

	if er != nil {
		fmt.Println("Error creating PUT request: ", er)
	}

	req.Header.Set("Content-type", "application/json")

	//Send the request
	client := http.Client{}
	res, ers := client.Do(req)
	if ers != nil {
		fmt.Println("Error sending request: ", ers)
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
}
func main() {
	performUpdateRequest()
}
