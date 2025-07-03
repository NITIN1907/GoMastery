package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task starting id: %d\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Task id ending: %d\n", id)

}
func main() {
	fmt.Println("practice sync-waitgroup")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	fmt.Println("ending task ids...")
	fmt.Println("Elapsed:", time.Since(start))
}
