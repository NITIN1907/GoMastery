package main

import "fmt"

func sendData(ch chan int) {
	fmt.Println("Sending value: ", ch)
	ch <- 42
}
func main() {
	ch := make(chan int)

	go sendData(ch)

	value := <-ch

	fmt.Println("Recieved: ", value)

}
