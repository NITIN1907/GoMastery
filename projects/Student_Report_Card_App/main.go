package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []float64
}

func (s Student) Average() float64 {
	var sum float64
	for _, i := range s.marks {
		sum += i
	}
	return sum / float64(len(s.marks))
}
func (s Student) Display() {
	fmt.Printf("Name: %s | roll no: %d\n", s.name, s.rollNo)
	fmt.Println("Marks: ", s.marks)
	fmt.Printf("Average: %.2f\n", s.Average())
}

func main() {
	var s Student
	s.name = "john"
	s.rollNo = 101
	s.marks = []float64{87.5, 98.4, 78.0}
	s.Display()
}
