package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

var mu sync.Mutex

//	func increment() {
//		for i := 1; i <= 10000000; i++ {
//			counter++
//		}
//	}
func increment() {
	for i := 1; i <= 10000000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
func main() {

	go increment()
	go increment()
	time.Sleep(time.Second)
	fmt.Println("final counter: ", counter)
}
