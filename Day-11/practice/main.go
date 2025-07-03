package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d running\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d done\n", id)
}
func greet(name string) {
	fmt.Println("hello : ", name)
}
func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}
	time.Sleep(2 * time.Second)

	go greet("goroutine")
	go greet("goroutine")
	go greet("goroutine")

}
