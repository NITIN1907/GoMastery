package main

import "fmt"

func main() {
	//A map is a built-in Go data type that stores key-value pairs (like dictionaries in Python).
	StudentGrade := make(map[string]int) // key is string, value is int

	StudentGrade["john"] = 80
	StudentGrade["jane"] = 92
	StudentGrade["kane"] = 89

	fmt.Println(StudentGrade["john"])

	//operation on map
	//1) add
	StudentGrade["tom"] = 95
	fmt.Println(StudentGrade)

	//2) delete
	delete(StudentGrade, "kane")
	fmt.Println(StudentGrade)

	//3)update
	StudentGrade["jane"] = 94
	fmt.Println(StudentGrade)

	//4)Check Key Exists
	value, exist := StudentGrade["john"] //value will store the actual value associated with the key "Bob", if it exists. exists is a boolean (true or false) that tells you whether the key "Bob" is present in the map.
	if exist {
		fmt.Println("Found: ", value)
	} else {
		fmt.Println("Not Found...")
	}
	//Zero value when key doesn't exist

	//Loop through map
	for name, marks := range StudentGrade {
		fmt.Printf("%s : %d\n", name, marks)
	}
}
