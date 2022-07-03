package main

import (
	"fmt"
)

func main() {
	x := 2

	if x == 2 {
		fmt.Printf("%v", x)
		lineVoid()
	}
}

func lineVoid() {
	fmt.Println()
}
