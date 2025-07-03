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
		time.Sleep(1 * time.Second)
		ch2 <- "hello from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Recieved: ", msg)
	case msg := <-ch2:
		fmt.Println("Recieved: ", msg)
	}

}
