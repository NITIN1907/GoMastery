package main

import "fmt"

func main() {
	var m1, m2, m3 float64
	fmt.Println("enter the marks: ")
	fmt.Scan(&m1, &m2, &m3)

	avg := (m1 + m2 + m3) / 2

	fmt.Println("Average: ", avg)

	if avg >= 90 {
		fmt.Println("Grade A")
	} else if avg >= 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

}
