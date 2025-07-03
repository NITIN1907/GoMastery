package main

import "fmt"

func main() {
	//declaring string(immutable nature)
	var st string = "Hello"
	st2 := "world"

	//concatenation
	result := st + " " + st2
	fmt.Println(result) //hello world

	//length
	fmt.Println(len(st))

	//character access
	s := "Go"
	fmt.Println(s[0])         //71 -->ASCII value
	fmt.Println(string(s[0])) //"G"

	//looping over a string
	for i, ch := range st {
		fmt.Printf("%d -> %c\n", i, ch)
	}

	//functions in string

}
