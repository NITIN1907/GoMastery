package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)

	fmt.Printf("worker %d end\n", i)

}
func main() {
	fmt.Println("Explore goroutine started...")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) //increment the waitgroup counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("workers task complete")
}
