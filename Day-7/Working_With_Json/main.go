package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("we are learning Json")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	// fmt.Println("Person Data is: ", person)

	//convert person into JSON Encoding(Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling....")
	}
	fmt.Println("person marshal Data is: ", string(jsonData)) //returns slice of bytes

	//Decoding (Unmarshalling)
	var personData Person
	er := json.Unmarshal(jsonData, &personData)
	if er != nil {
		fmt.Println("Error unmarshalling", er)
		return
	}
	fmt.Println("Person unmarshal Data is : ", personData)
}
