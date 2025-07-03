package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"
)

func main() {
	//web request
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get response", err)
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res) //Type of response: *http.Response
	//fmt.Println(res)

	//Read the reponses Body

	data, er := io.ReadAll(res.Body)
	if er != nil {
		fmt.Println("error...", er)
		return
	}
	fmt.Println(string(data))
}

//In Go, you can make web requests using net/http package, which provides funtions to create and send HTTP requests, as well as handle responses
