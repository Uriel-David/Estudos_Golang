package main

import (
	"fmt"
)

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 89}
	fmt.Println(function1(si...))
	fmt.Println(function2(si))
}

func function1(x ...int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func function2(x []int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func lineVoid() {
	fmt.Println()
}
