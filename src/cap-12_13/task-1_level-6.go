package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d", givesItBackInt())
	lineVoid()
	fmt.Println(givesItBackIntAndString())
}

func givesItBackInt() int {
	return 10
}

func givesItBackIntAndString() (int, string) {
	return 20, "vinte"
}

func lineVoid() {
	fmt.Println()
}
