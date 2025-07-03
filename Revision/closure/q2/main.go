package main

import "fmt"

func makeFibo() func(int) int {

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	return fib
}
func main() {
	fmt.Println("hello")
	fibo := makeFibo()
	fmt.Println(fibo(5))
}
