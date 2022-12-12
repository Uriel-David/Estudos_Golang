package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (q square) area() {
	result := q.side * q.side
	fmt.Println("Square area is:", result)
}

type circle struct {
	radius float64
}

func (c circle) area() {
	result := math.Pi * 2 * c.radius
	fmt.Println("Circle area is:", result)
}

type info interface {
	area()
}

func metric(i info) {
	i.area()
}

func main() {

	x := square{
		side: 15.0,
	}

	y := circle{
		radius: 5.25,
	}

	//x.area()
	//y.area()

	metric(x)
	metric(y)
}

func lineVoid() {
	fmt.Println()
}
