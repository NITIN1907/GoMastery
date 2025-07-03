package main

import (
	"fmt"
	// "time"
)

// var count int // Shared data

func increment(ch chan int) {
	count := 0
	for i := 0; i < 1000000; i++ {
		count++
	}
	ch <- count
}

func main() {
	ch := make(chan int)
	go increment(ch)
	go increment(ch)
	count1 := <-ch
	count2 := <-ch
	// go increment()
	// go increment()
	// time.Sleep(1 * time.Second) // Give time for goroutines to finish
	fmt.Println("final count: ", count1+count2) //always show accurate data(20000000)
	// fmt.Println("Final Count:", count)//not same different
}
