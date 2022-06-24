package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	lineVoid()

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)
	fmt.Printf("z: %v %T\n", z, z)
}

func lineVoid() {
	fmt.Println()
}
