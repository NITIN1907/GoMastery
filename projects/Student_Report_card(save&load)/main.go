package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name   string    `json:"name"`
	Rollno int       `json:"rollno"`
	Marks  []float64 `json:"marks"`
}

func main() {
	//save to file
	student := Student{"john", 101, []float64{86.8, 98.0, 87.8, 95.5}}
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("Error marshalling...")
	}
	os.WriteFile("report.json", data, 0644)
	fmt.Println("Report saved to file....")

	//load from file
	readData, er := os.ReadFile("report.json")
	if er != nil {
		fmt.Println("Error....")
	}
	var load Student
	json.Unmarshal(readData, &load)
	fmt.Println("Loaded Student: ", load)
}
