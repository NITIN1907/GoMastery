package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("helloo")
	// time.Sleep(2000 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello after sleep")
}
func sayGreet() {
	fmt.Println("greet")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("greeting after sleep")
}
func main() {
	fmt.Println("learning gorountine")
	go sayHello()
	go sayGreet()
	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(2000 * time.Millisecond)
}
