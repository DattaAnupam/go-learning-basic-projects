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

// Definition for interface's function for square struct
func (s square) area() float64 {
	return s.side * s.side
}

// Prints the area
func printArea(s shape) {
	fmt.Println(s.area())
}
