package main

import "fmt"

func main() {
	student := make(map[string]int)

	for {
		fmt.Println("Enter your choice: \n1) student detail \n2) average \n3) exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("enter the student name: ")
			var name string
			fmt.Scan(&name)
			fmt.Println("Enter the marks of student: ")
			var marks int
			fmt.Scan(&marks)
			student[name] = marks
			fmt.Println("Details added..")
		case 2:
			var avg float64
			var sum int
			for _, i := range student {
				sum += i
			}
			avg = float64(sum) / float64(len(student))
			fmt.Println("Average: ", avg)
		case 3:
			fmt.Println("Existing.....")
			return
		}

	}
}
