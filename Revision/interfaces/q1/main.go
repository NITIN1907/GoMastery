package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}
func main() {
	r := Rectangle{
		length:  12.4,
		breadth: 14.5,
	}
	c := Circle{
		radius: 2.4,
	}

	PrintArea(r)
	PrintArea(c)
}
