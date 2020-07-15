package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// interfaces
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}
