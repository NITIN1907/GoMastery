package main

import (
	"fmt"
	"sync"
)

func sum(num1 []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, add := range num1 {
		sum += add
	}
	result <- sum
}
func main() {
	num := []int{1, 2, 3, 4, 5}

	mid := len(num) / 2
	firsthalf := num[:mid]
	secondhalf := num[mid:]

	result := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go sum(firsthalf, result, &wg)
	go sum(secondhalf, result, &wg)

	wg.Wait()
	close(result)

	total := 0
	for ans := range result {
		total += ans
	}

	fmt.Println("Done", total)
}
