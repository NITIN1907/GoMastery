package main

import (
	"fmt"
	"time"
)

func main() {
	ch2 := make(chan int)
	ch1 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()
	go func() {
		ch3 <- 3
	}()

	total := 0
	for {
		select {
		case ans := <-ch1:
			fmt.Println("recieved: ", ans)
			total += ans
		case ans := <-ch2:
			fmt.Println("Got:", ans)
			total += ans
		case ans := <-ch3:
			fmt.Println("Got: ", ans)
			total += ans
		case <-time.After(5 * time.Second):
			fmt.Println("No activity for 5 seconds. Exiting.")
			return
		}
	}
	fmt.Println(total)
}
