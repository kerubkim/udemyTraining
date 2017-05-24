package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

// Both type Square and Circle implement Shape interface
type Shape interface {
	area() float64 // Declaration of method area
}

// When area is being called in from a Square. This area runs.
func (s Square) area() float64 { //* area() float64 -> needs to match what was Declared as Area in Shape
	return s.side * s.side
}

// When area is being called in from a Circle. This area runs.
func (c Circle) area() float64 { //* area() float64 -> needs to match what was Declared as Area in Shape
	return math.Pi * c.radius * c.radius
}

// since info takes in a type Shape, we can pass in a Square or a Circle
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	square := Square{10}
	circle := Circle{5}

	info(square)
	info(circle)

}
