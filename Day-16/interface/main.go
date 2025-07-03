package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "woof"
}
func (c Cat) Speak() string {
	return "meow!"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}
func PrintAnythin(i interface{}) {
	fmt.Println(i)
}
func main() {
	dog := Dog{}
	cat := Cat{}

	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)

	PrintAnythin(1)
	PrintAnythin("hello")
}
