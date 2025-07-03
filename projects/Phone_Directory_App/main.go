package main

import "fmt"

func main() {
	phoneBook := make(map[string]string)

	for {
		fmt.Println("Enter your choice: \n1. Add Contact\n2. Search Contact\n3. Delete Contact\n4. Exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the name: ")
			var name string
			fmt.Scan(&name)
			fmt.Println("Enter the phone number: ")
			var num string
			fmt.Scan(&num)
			phoneBook[name] = num
			fmt.Println("Contact saved.")
		case 2:
			fmt.Println("Enter the name to search: ")
			var name string
			fmt.Scan(&name)
			num, exist := phoneBook[name]
			if exist {
				fmt.Println("Number: ", num)
			} else {
				fmt.Println("Contact does not exist.")
			}
		case 3:
			fmt.Println("Enter the name to delete: ")
			var name string
			fmt.Scan(&name)
			delete(phoneBook, name)
			fmt.Println("Contact deleted if it existed.")
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
