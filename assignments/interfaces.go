package main

import "fmt"

type shape interface {
	printArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (s square) printArea() float64 {
	area := s.sideLength * s.sideLength

	fmt.Printf("Sqaure Area : %f\n", area)

	return area
}

func (t triangle) printArea() float64 {
	area := 0.5 * t.base * t.height

	fmt.Printf("Triangle Area : %f\n", area)

	return area
}

func main() {

	triangle := triangle{
		height: 5.5,
		base:   1,
	}

	triangle.printArea()

	square := square{
		sideLength: 10,
	}

	square.printArea()
}
