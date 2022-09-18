package main

import (
	"fmt"
)

type veichule struct {
	ports int
	color string
}

type pickup struct {
	veichule
	tractionFourWheels bool
}

type sedan struct {
	veichule
	modelLux bool
}

func main() {
	carOfUncle := sedan{veichule{4, "abobora"}, true}
	carOfGrandad := pickup{
		veichule: veichule{
			ports: 8,
			color: "rust",
		},
		tractionFourWheels: false,
	}

	fmt.Println(carOfUncle)
	fmt.Println(carOfGrandad)

	lineVoid()

	fmt.Println(carOfUncle.color)
	fmt.Println(carOfGrandad.tractionFourWheels)
}

func lineVoid() {
	fmt.Println()
}
