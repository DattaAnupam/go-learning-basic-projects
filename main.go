package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	// every struct who wants to implement this interface
	// should have a definition of all the functions defined
	// inside the interface
	area() float64
	perimeter() float64
}

func main() {
	t := triangle{
		base:   2.1,
		height: 3,
	}
	s := square{
		side: 3,
	}

	// print area of triangle
	printArea(t)
	// print area of square
	printArea(s)
}

// Definition for interface's function for triangle struct
func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
func (t triangle) perimeter() float64 {
	// considering Equilateral triangle
	return 3 * t.base
}

// Definition for interface's function for square struct
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

// Prints the area
func printArea(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}
