package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hello from ch2"
	}()

	// select {
	// case msg := <-ch1:
	// 	fmt.Println("Got: ", msg)
	// case msg := <-ch2:
	// 	fmt.Println("Got: ", msg)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("TimeOut!")
	// }
	for {
		select {
		case msg := <-ch1:
			fmt.Println("ch1:", msg)
		case msg := <-ch2:
			fmt.Println("ch2:", msg)
		case <-time.After(5 * time.Second):
			fmt.Println("No activity for 5 seconds. Exiting.")
			return
		}
	}

}
