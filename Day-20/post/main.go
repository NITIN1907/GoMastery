package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// "io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserId:    21,
		Id:        3,
		Title:     "Go lang",
		Completed: true,
	}
	myurl := "https://jsonplaceholder.typicode.com/todos"

	//convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	//convert the json byte slice to a string
	jsonString := string(jsonData)

	// fmt.Println("jsonString: ", jsonString)

	//create an io.Reader from the string
	jsonReader := strings.NewReader(jsonString)

	// fmt.Println("jsonReader", jsonReader)

	//Send the Post request
	res, er := http.Post(myurl, "application/json", jsonReader)
	if er != nil {
		fmt.Println("Error sending request: ", er)
		return
	}

	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response status: ", res.Status)
	fmt.Println("Response: ", string(data))
}
func main() {
	fmt.Println("Learning crud...")
	// performGetRequest()
	performPostRequest()
	performGetRequest()
}

//decode JSON data from the HTTP response body and convert it directly into a Go struct (Todo in your case).
//This function creates a new JSON decoder that reads from an io.Reader stream (like res.Body in this case).

//It doesn’t read all data at once — it reads incrementally, which is more memory-efficient for large JSON.
