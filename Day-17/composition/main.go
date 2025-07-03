package main

import "fmt"

type Engine struct {
	HorsePower int
}
type Car struct {
	Engine
	Brand string
}

func (c Car) speed() {
	fmt.Println("brand name: ", c.Brand)
	fmt.Println("Top speed: ", c.Engine.HorsePower, "kmph")
}
func main() {
	c := Car{
		Brand: "Lambo",
		Engine: Engine{
			HorsePower: 423,
		},
	}
	fmt.Println(c.Brand)
	c.speed()

}
