package main

import "fmt"

func main() {
	fmt.Println("if_else")
	score := 45
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 70 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	//switch case

	var greet string
	fmt.Println("greet me: ")
	fmt.Scan(&greet)
	switch greet {
	case "hello":
		fmt.Println("hey, how are you?")
	case "hey":
		fmt.Println("hello, how are you?")
	default:
		fmt.Println("Good morning......")
	}

	//for loop
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

}
