package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x
	fmt.Println(&x)
	fmt.Println(*&x)
	fmt.Println(&y)
	fmt.Println(*y)
	fmt.Printf("%T, %T", x, y)

	lineVoid()

	*y++
	fmt.Println(x, y)

	lineVoid()

	z := 0
	receiveValue(z)
	fmt.Println(z)

	receivePointer(&z)
	fmt.Println(z)
}

func receiveValue(x int) {
	x++
	fmt.Println("Passed-Value: ", x)
}

func receivePointer(z *int) {
	*z++
	fmt.Println("Pointer: ", *z)
}

func lineVoid() {
	fmt.Println()
}
