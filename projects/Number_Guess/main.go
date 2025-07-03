package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(100) + 1

	var guess int
	for i := 0; i <= 10; i++ {
		fmt.Println("guess the number: ")
		fmt.Scan(&guess)

		if guess == secret {
			fmt.Println("🎉 Correct! You win!")
			return
		} else if guess < secret {
			fmt.Println("guess is low")
		} else {
			fmt.Println("guess is high")
		}
	}
	fmt.Println("💥 Out of attempts! The number was:", secret)
}
