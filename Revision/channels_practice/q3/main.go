package main

import "fmt"

func sum(num []int, result chan int) {
	sum := 0
	for _, ans := range num {
		sum += ans
	}
	result <- sum
}
func main() {
	num := []int{1, 2, 3, 4, 5}

	mid := len(num) / 2
	firsthalf := num[:mid]
	secondhalf := num[mid:]

	result := make(chan int)

	go sum(firsthalf, result)
	go sum(secondhalf, result)

	sum1 := <-result
	sum2 := <-result

	fmt.Println(sum1)
	fmt.Println(sum2)

	total := sum1 + sum2
	fmt.Println(total)
}
