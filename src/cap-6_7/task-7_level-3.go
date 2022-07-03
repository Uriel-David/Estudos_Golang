package main

import (
	"fmt"
)

func main() {
	x := 1

	if x == 2 {
		fmt.Printf("%v", x)
		lineVoid()
	} else if x == 1 {
		fmt.Printf("%v", x)
		lineVoid()
	} else {
		fmt.Printf("%v", x)
		lineVoid()
	}
}

func lineVoid() {
	fmt.Println()
}
