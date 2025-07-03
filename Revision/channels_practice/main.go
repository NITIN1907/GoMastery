package main

import "fmt"

func main() {
	fmt.Println()
	counter := 0
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 1; i <= 2000; i++ {
			counter += <-ch
		}
		done <- true
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- 1
		}
	}()
	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- 1
		}
	}()

	<-done
	fmt.Println("counter: ", counter)
}
