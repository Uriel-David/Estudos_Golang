package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

var h = 10

type hotdog int

var newHotdog hotdog

func main() {
	x := "String"
	y := 10
	z := true
	w := 10.0

	newHotdog = 200

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)
	fmt.Printf("z: %v %T\n", z, z)
	fmt.Printf("w: %v %T\n\n", w, w)

	y = 20 * 2
	fmt.Printf("y: %v %T\n\n", y, y)

	z = 10 != 10
	fmt.Printf("z: %v %T\n\n", z, z)

	fmt.Println(x, y, z, w)

	firstVar, errors := fmt.Println("Hello world!", "Hey Guys...", 100)
	fmt.Println(firstVar, errors)

	lineVoid()

	qualquerCoisa(y)

	valorZero()

	print := fmt.Sprint(x, " = ", x)

	fmt.Println(print)

	fmt.Printf("newHotdog: %v %T\n\n", newHotdog, newHotdog)

	y = int(newHotdog) // y = newHotdog -> problem types diff
	fmt.Printf("y: %v %T\n", y, y)
}

func qualquerCoisa(x int) {
	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("h: %v %T\n\n", h, h)
}

func valorZero() {
	fmt.Printf("a: %v %T\n", a, a)
	fmt.Printf("b: %v %T\n", b, b)
	fmt.Printf("c: %v %T\n", c, c)
	fmt.Printf("d: %v %T\n\n", d, d)
}

func lineVoid() {
	fmt.Println()
}
