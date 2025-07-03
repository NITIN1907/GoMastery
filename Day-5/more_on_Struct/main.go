package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House string
	Area  string
	State string
}

type Employee struct {
	Person_Detail  Person
	Person_Contact Contact
	Person_Address Address
}

func (e Employee) Display() {
	fmt.Println("Employee Detail", e.Person_Detail, "| Employee contact: ", e.Person_Contact)
}
func main() {
	var emp1 Employee
	emp1.Person_Detail = Person{
		firstName: "Jane",
		lastName:  "Watson",
		age:       24,
	}
	emp1.Person_Contact.Email = "contact@gmail.com"
	emp1.Person_Contact.Phone = "984838283"

	emp1.Person_Address = Address{
		House: "Green Homes",
		Area:  "New Delhi",
		State: "Delhi",
	}

	fmt.Println(emp1)                //{{Jane Watson 24} {contact@gmail.com 984838283} {Green Homes New Delhi Delhi}}
	fmt.Println(emp1.Person_Address) //{Green Homes New Delhi Delhi}
	emp1.Display()
}
