package main

import "fmt"

func main() {
	arr := []string{"apple", "banana"}
	var option string
	for {
		fmt.Print("Enter to add or delete the array: ")
		fmt.Scan(&option)

		switch option {
		case "add":
			var name string
			fmt.Println("Enter the name: ")
			fmt.Scan(&name)
			arr = append(arr, name)
		case "delete":
			var i int
			fmt.Println("choose the index to delete: ")
			fmt.Scan(&i)
			arr = append(arr[:i], arr[i+1:]...)
		case "display":
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i])
			}
		default:
			fmt.Print("exiting.....")
			return
		}
	}
}
