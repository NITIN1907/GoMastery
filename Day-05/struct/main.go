package main

import "fmt"

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

type Student struct {
	name  string
	roll  int
	marks float64
}

func main() {
	//?A struct is a composite data type (also called an aggregate) in Go. It groups together variables (called fields) under one name. Structs are used to represent real-world entities with different attributes. It's similar to class in other object-oriented languages (but without inheritance).

	//1st method Declare and Initialize

	// var first Person
	// fmt.Println("First Person: ", first)
	// first.FirstName = "sample"
	// first.LastName = "lastSample"
	// first.Age = 24
	// fmt.Println("First Person: ", first)

	//2nd method With Composite Literal
	s := Student{name: "sample", roll: 101, marks: 98.5}

	fmt.Println(s.name, s.roll, s.marks)

	//3rd method using new keyword
	var std1 = new(Student)
	std1.name = "Emma"
	std1.roll = 02
	std1.marks = 94.4
	fmt.Println(std1) //&{sample 101 98.5} because `new(Student)` allocates memory for a Student struct and returns a pointer to it.
	// std1 is a *Student, and when printed, Go automatically dereferences the pointer to show the field values.
	fmt.Println(std1.name) //sample

}
